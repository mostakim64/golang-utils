package logger

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/shadowchef/utils/slackit"
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

func ProcessAndSend(slackLogReq SlacklogRequest, status int, logType string) error {

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
				return fmt.Errorf("Failed while sending to slack webhook [%v]", err)
			}
		}
	}

	return nil
}

func ProcessAndSendWithMeta(slackLogReq SlacklogRequest, metaData interface{}, status int, logType string) error {

	if slackitClient != nil {

		metaJson, err := json.MarshalIndent(metaData, "", "  ")
		if err != nil {
			return err
		}

		msg, err := json.MarshalIndent(&slackLogReq, "", "\t")
		if err != nil {
			return err
		}
		if msg != nil {
			clientReq := slackit.ClientRequest{
				Header:      slackLogReq.Level,
				ServiceName: serviceName,
				Summary:     logType + " Log from " + serviceName,
				Metadata:    string(metaJson),
				Details:     string(msg),
				Status:      status,
			}
			err = slackitClient.Send(clientReq)
			if err != nil {
				return fmt.Errorf("Failed while sending to slack webhook [%v]", err)
			}
		}
	}

	return nil
}

func ProcessAndSendWitResAndMeta(slackLogReq SlacklogRequestWithRes, metaData interface{}, status int, logType string) error {

	if slackitClient != nil {

		metaJson, err := json.MarshalIndent(metaData, "", "  ")
		if err != nil {
			return err
		}

		msg, err := json.MarshalIndent(&slackLogReq, "", "\t")
		if err != nil {
			return err
		}

		var mentions []string
		if status == slackit.Alert {
			mentions = append(mentions, "@here")
		}

		if msg != nil {
			clientReq := slackit.ClientRequest{
				Header:      "API " + slackLogReq.Level,
				ServiceName: serviceName,
				Summary:     logType + " Log from " + serviceName,
				Metadata:    string(metaJson),
				Details:     string(msg),
				Status:      status,
				Mentions:    mentions,
			}
			err = slackitClient.Send(clientReq)
			if err != nil {
				return fmt.Errorf("Failed while sending to slack webhook [%v]", err)
			}
		}
	}

	return nil
}
