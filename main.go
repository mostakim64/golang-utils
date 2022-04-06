package main

import (
	"bitbucket.org/shadowchef/utils/logger"
	"errors"
)

func main() {
	webhookUrl := "https://hooks.slack.com/services/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"
	service := "storage"
	_ = logger.NewSlackLogitClient(webhookUrl, service)
	e := errors.New("new error")
	logger.SetLogJsonFormatter()
	logger.Error("Error occurred", e.Error(), nil)
	logger.Info("Test Info ", e.Error(), nil)
	logger.Debug("Test debug")

}