package translation

import (
	"runtime"

	"github.com/klikit/utils/logger"
)

func GetCallerFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

type MissingTranslationMeta struct {
	Caller   string `json:"caller"`
	Language string `json:"language"`
}

type MissingTranslationInfo struct {
	Message string   `json:"message,omitempty"`
	Error   string   `json:"error,omitempty"`
	Fields  []string `json:"fields,omitempty"`
}

func (MissingTranslationInfo) GetLevel() string {
	return "info"
}

func sendNotification(
	caller string,
	lang string,
	fields []string,
	err string,
) {

	var message string
	if len(fields) > 0 {
		message = "Some fields are missing translations for validation errors"
	} else if err != "" {
		message = "Translation not found for the error"
	}

	logger.PushSlack(
		MissingTranslationMeta{
			Caller:   caller,
			Language: lang,
		},
		MissingTranslationInfo{
			Message: message,
			Error:   err,
			Fields:  fields,
		},
	)
}
