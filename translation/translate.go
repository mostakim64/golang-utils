package translation

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/klikit/utils/slackit"
	"sync"
)

var slackitClient slackit.SlackitClient
var callerService string

func InitLogger(slackURL string, service string) {
	slackitClient = slackit.NewSlackitClient(slackURL)
	callerService = service
}

func MergeMapper(mapper map[string]map[string]string) {
	for key, value := range mapper {
		translations[key] = value
	}
}

func TranslateError(err error, lang string) error {
	var validationErrors validation.Errors
	translatedErrors := make(validation.Errors)

	missingFields := make(map[string]string)
	var missingTranslation string

	caller := GetCallerFuncName()
	defer func() {
		if len(missingFields) > 0 || missingTranslation != "" {
			sendNotification(caller, lang, missingFields, missingTranslation)
		}
	}()

	if errors.As(err, &validationErrors) {
		for field, validationErr := range validationErrors {
			translatedErrors[field] = translateNestedError(validationErr, lang, missingFields)
		}
		return translatedErrors
	}

	// If the error is not validation-specific, track missing translations for this case
	missingTranslation = err.Error()

	return err
}

// Helper function to recursively translate nested errors
func translateNestedError(err error, lang string, missingFields map[string]string) error {
	var ve validation.Error
	var nestedErrors validation.Errors

	// If the error is another map of validation errors, recurse into it
	if errors.As(err, &nestedErrors) {
		translatedNestedErrors := make(validation.Errors)
		var firstKey string
		var once sync.Once
		for nestedField, nestedErr := range nestedErrors {
			// Recurse into nested errors
			translatedNestedErrors[nestedField] = translateNestedError(nestedErr, lang, missingFields)
			once.Do(func() {
				firstKey = nestedField
			})
		}
		return translatedNestedErrors[firstKey]
	}

	// Translate individual validation error
	if errors.As(err, &ve) {
		if group := translations[ve.Code()]; len(group) > 0 {
			if trans, ok := group[lang]; ok {
				return ve.SetMessage(trans) // Return translated message
			}
		}
		// If no translation, add to missingFields
		missingFields[ve.Code()] = ve.Error()
		return ve // Return untranslated error
	}

	// If it's not a validation error and no further nesting is possible, just return the error
	missingFields["general"] = err.Error()
	return err
}
