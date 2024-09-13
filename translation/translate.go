package translation

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func TranslateError(err error, lang string) error {
	var validationErrors validation.Errors
	translatedErrors := make(validation.Errors)

	var missingFields []string
	var missingTranslation string

	caller := GetCallerFuncName()
	defer func() {
		sendNotification(caller, lang, missingFields, missingTranslation)
	}()

	if errors.As(err, &validationErrors) {
		for field, validationErr := range validationErrors {
			var ve validation.Error
			if errors.As(validationErr, &ve) {
				if group := translations[ve.Code()]; len(group) > 0 {
					if trans, ok := group[lang]; ok {
						vErr := ve.SetMessage(trans)
						translatedErrors[field] = vErr
						continue
					}
				}
			}
			missingFields = append(missingFields, field)
			translatedErrors[field] = validationErr
		}
		return translatedErrors
	}
	missingTranslation = err.Error()

	return err
}
