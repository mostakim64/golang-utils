package main
import (

	"bitbucket.org/shadowchef/utils/logger"
)

func main() {
	webhookUrl := "https://hooks.slack.com/services/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"
	service := "storage"
	_ = logger.NewSlackLogitClient(webhookUrl, service)
	logger.Error("Error occurred", "Error 2", "Error 3")
}