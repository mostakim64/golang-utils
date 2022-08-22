package logger

import "github.com/sirupsen/logrus"

// fields wraps logrus.Fields, which is a map[string]interface{}
type fields logrus.Fields

type SlacklogRequest struct {
	Message string `json:"message"`
	File    string `json:"file"`
	Level   string `json:"level"`
}

type KlikitLogger struct {
	client *logrus.Logger
}
