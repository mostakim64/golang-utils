package logger

import (
	"bitbucket.org/shadowchef/utils/slackit"
	"encoding/json"
)

var slackitClient *slackit.SlackitClient
var serviceName string

func SetSlackLogger(webhookUrl, service string) {
	client := slackit.NewSlackitClient(webhookUrl)
	slackitClient = &client
	serviceName = service
}

func send(msg string) {
	clientReq := slackit.ClientRequest{
		Header:      "Alert",
		ServiceName: serviceName,
		Summary:     "Some error occurred TEST",
		Details:     msg,
		Status:      slackit.Alert,
	}

	_ = slackitClient.Send(clientReq)
}

func ProcessAndSend(slackLogReq SlacklogRequest, status int, logType string) error{

	if slackitClient != nil {
		msg, err := json.MarshalIndent(&slackLogReq, "", "\t")
		if err != nil {
			return err
		}
		if msg != nil {
			clientReq := slackit.ClientRequest{
				Header:      slackLogReq.Level,
				ServiceName: serviceName,
				Summary:     logType + " Log from " + serviceName,
				Details:     string(msg),
				Status:      status,
			}
			err = slackitClient.Send(clientReq)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
