package logger

import (
	"bitbucket.org/shadowchef/utils/slackit"
	"encoding/json"
)

var slackitClient *slackit.SlackitClient
var serviceName string

func NewSlackLogitClient(webhookUrl, service string) error {
	client := slackit.NewSlackitClient(webhookUrl)
	slackitClient = &client
	serviceName = service
	return nil
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

func ProcessAndSend(slackLogReq SlacklogRequest, status int) error{

	if slackitClient != nil {
		msg, err := json.MarshalIndent(&slackLogReq, "", "\t")
		if err != nil {
			return err
		}
		if msg != nil {
			clientReq := slackit.ClientRequest{
				Header:      slackLogReq.Level,
				ServiceName: serviceName,
				Summary:     "Some error occurred TEST",
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
