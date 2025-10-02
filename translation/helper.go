package translation

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/mostakim64/golang-utils/slackit"
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
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func sendNotification(
	caller string,
	lang string,
	fields map[string]string,
	err string,
) {

	var message string
	if len(fields) > 0 {
		message = "Some fields are missing translations for validation errors"
	} else if err != "" {
		message = "Translation not found for the error"
	}

	metadata := MissingTranslationMeta{
		Caller:   caller,
		Language: lang,
	}
	metaByte, _ := json.MarshalIndent(metadata, "", " ")

	details := MissingTranslationInfo{
		Message: message,
		Error:   err,
		Fields:  fields,
	}
	detailsByte, _ := json.MarshalIndent(details, "", " ")

	request := slackit.ClientRequest{
		Header:      fmt.Sprintf("Missing translation in %s", callerService),
		ServiceName: callerService,
		Summary:     message,
		Metadata:    string(metaByte),
		Details:     string(detailsByte),
		Status:      slackit.Warning,
	}

	_ = slackitClient.Send(request)
}
