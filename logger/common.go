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

func NewLoggerClient() *KlikitLogger {
	return &KlikitLogger{
		client: logrus.New(),
	}
}

func (r *KlikitLogger) SetLogLevel(level logrus.Level) {
	r.client.Level = level
}

func (r *KlikitLogger) SetLogFormatter(formatter logrus.Formatter) {
	r.client.Formatter = formatter
}

func (r *KlikitLogger) SetLogJsonFormatter() {
	r.client.Formatter = &logrus.JSONFormatter{}
}

// Debug logs a message at level Debug on the KlikitLogger.
func (r *KlikitLogger) Debug(args ...interface{}) {
	if r.client.Level >= logrus.DebugLevel {
		entry := r.client.WithFields(logrus.Fields{})
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// DebugWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) DebugWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.DebugLevel {
		entry := r.client.WithFields(f)
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the KlikitLogger.
func (r *KlikitLogger) Info(args ...interface{}) {
	if r.client.Level >= logrus.InfoLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// InfoWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) InfoWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.InfoLevel {
		entry := r.client.WithFields(f)
		//entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the KlikitLogger.
func (r *KlikitLogger) Warn(args ...interface{}) {
	if r.client.Level >= logrus.WarnLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) WarnWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.WarnLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the KlikitLogger.
//
// if 1 item in args then there will be no metadata
//
// if multiple items in args then 1st item will be treated as metadata and rest items will go for args
func (r *KlikitLogger) Error(args ...interface{}) {
	var metaData interface{}

	if len(args) > 1 {
		metaData = args[0]
		args = args[1:]
	}

	if r.client.Level >= logrus.ErrorLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)

		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "error",
		}
		if err := ProcessAndSendWithMeta(slackLogReq, metaData, slackit.Alert, "Error"); err != nil {
			r.Warn(err)
		}
	}
}

// Error logs a message at level Error on the KlikitLogger. with res data and metaData
func (r *KlikitLogger) ApiError(rs RequestResponseMap, metaData interface{}, args ...interface{}) {
	if r.client.Level >= logrus.ErrorLevel {
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
			r.Warn(err)
		}
	}
}

// ErrorWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) ErrorWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.ErrorLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the KlikitLogger.
func (r *KlikitLogger) Fatal(args ...interface{}) {
	if r.client.Level >= logrus.FatalLevel {
		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "fatal",
		}
		_ = ProcessAndSend(slackLogReq, slackit.Alert, "Fatal")
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)

	}
}

// FatalWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) FatalWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.FatalLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the KlikitLogger.
func (r *KlikitLogger) Panic(args ...interface{}) {
	if r.client.Level >= logrus.PanicLevel {

		slackLogReq := SlacklogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "panic",
		}
		_ = ProcessAndSend(slackLogReq, slackit.Alert, "Panic")

		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)

	}
}

// PanicWithFields Debug logs a message with fields at level Debug on the KlikitLogger.
func (r *KlikitLogger) PanicWithFields(l interface{}, f fields) {
	if r.client.Level >= logrus.PanicLevel {
		entry := r.client.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func (r *KlikitLogger) fileInfo(skip int) string {
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

func (r *KlikitLogger) fileAddressInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file, line)
}

func (r *KlikitLogger) processLog(args ...interface{}) string {
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
