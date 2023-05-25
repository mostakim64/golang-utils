package logger

import (
	"net/http"

	"github.com/klikit/utils/slackit"
	"github.com/sirupsen/logrus"
)

// fields wraps logrus.Fields, which is a map[string]interface{}
type fields logrus.Fields

type SlacklogRequest struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Level   string `json:"level"`
}

type SlacklogRequestWithApiError struct {
	Message    string           `json:"message"`
	File       string           `json:"file"`
	Level      string           `json:"level"`
	ApiDetails slackit.ApiError `json:"api_details"`
}

type KlikitLogger struct {
	client *logrus.Logger
}

type RequestResponseMap struct {
	Req     *http.Request
	ReqBody interface{}
	Res     *http.Response
	ResBody interface{}
}
