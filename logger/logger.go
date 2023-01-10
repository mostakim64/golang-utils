package logger

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"bitbucket.org/shadowchef/utils/slackit"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

func SetLogJsonFormatter() {
	logger.Formatter = &logrus.JSONFormatter{}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// DebugWithFields Debug logs a message with fields at level Debug on the standard logger.
func DebugWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// InfoWithFields Debug logs a message with fields at level Debug on the standard logger.
func InfoWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		//entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields Debug logs a message with fields at level Debug on the standard logger.
func WarnWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger.
//
// if 1 item in args then there will be no metadata
//
// if multiple items in args then 1st item will be treated as metadata and rest items will go for args
func Error(args ...interface{}) {
	var metaData interface{}

	if len(args) > 1 {
		metaData = args[0]
		args = args[1:]
	}

	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "error",
		}
		if err := ProcessAndSendWithMeta(slackLogReq, metaData, slackit.Alert, "Error"); err != nil {
			Warn(err)
		}
	}
}

// Error logs a message at level Error on the standard logger with request, response and metadata
func ApiError(rs RequestResponseMap, metaData interface{}, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
		whichApi := args[0].(string)

		slackLogReq := SlacklogRequestWithApiError{
			Message: fmt.Sprint(args...) + " Failed",
			File:    fileAddressInfo(2),
			Level:   "error",
			ApiDetails: slackit.ApiError{
				Api: whichApi,
				Url: rs.Req.URL.String(),
				Data: slackit.Data{
					Status:       rs.Res.StatusCode,
					Headers:      rs.Req.Header,
					RequestBody:  rs.ReqBody,
					ResponseBody: rs.ResBody,
				},
			},
		}

		if err := ProcessAndSendWithApiError(slackLogReq, metaData, slackit.Alert, "Error"); err != nil {
			Warn(err)
		}

	}
}

// ErrorWithFields Debug logs a message with fields at level Debug on the standard logger.
func ErrorWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "fatal",
		}
		_ = ProcessAndSend(slackLogReq, slackit.Alert, "Fatal")
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)

	}
}

// FatalWithFields Debug logs a message with fields at level Debug on the standard logger.
func FatalWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {

		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "panic",
		}
		_ = ProcessAndSend(slackLogReq, slackit.Alert, "Panic")

		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)

	}
}

// PanicWithFields Debug logs a message with fields at level Debug on the standard logger.
func PanicWithFields(l interface{}, f fields) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func fileAddressInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file, line)
}

func processLog(args ...interface{}) string {
	var errMsgBuffer bytes.Buffer
	for _, arg := range args {
		if arg != nil {
			switch reflect.TypeOf(arg).Kind() {
			case reflect.String:
				errMsgBuffer.WriteString(arg.(string) + "\n")
			case reflect.Ptr:
				e := arg.(error)
				errMsgBuffer.WriteString(e.Error() + "\n")
			}
		}
	}

	return errMsgBuffer.String()
}
