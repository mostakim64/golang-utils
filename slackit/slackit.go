package slackit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var client slackitClient


type slackitClient struct {
	webhookUrl string
}

func Slackit() slackitClient {
	return client
}

func NewSlackitClient(webhookUrl string) slackitClient{
	client = slackitClient{
		webhookUrl: webhookUrl,
	}
	return client
}

func generateSingleBlock(typ string, text *Text) Blocks {
	block := Blocks{
		Type: typ,
		Text: text,
	}
	return block
}

func generateSectionBlock(fields []*Fields) Blocks {
	block := Blocks{
		Type: "section",
		Fields: fields,
	}
	return block
}

func generateText(typ string, txt string, emoji *bool) *Text{
	text := Text{
		Type: typ,
		Text: txt,
		Emoji: emoji,
	}
	return &text
}

func generateFields(typ string, txt string) *Fields {
	field := Fields{
		Type: typ,
		Text: txt,
	}
	return &field
}

func (sc *slackitClient) SendSlackNotification(serviceName string , summary string, details string) error {

	currentTime := time.Now()

	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")

	emoji := true

	headerText := generateText("plain_text", "New Alert", &emoji)

	headerBlock := generateSingleBlock("header", headerText)

	serviceNameField := generateFields("mrkdwn", "*Service:*\nStorage")

	serviceLogTimeField := generateFields("mrkdwn", "*Created At:*\n"+currentTimeStr)

	serviceInfoBlock := generateSectionBlock([]*Fields{serviceNameField, serviceLogTimeField})

	summaryField := generateFields("mrkdwn", "*Summary:*\n"+summary)

	summaryBlock := generateSectionBlock([]*Fields{summaryField})

	var detailsBlocks []Blocks
	detailsArr := Chunks(details, 1000)

	for ind, detail := range detailsArr{
		//detailsField := generateFields("mrkdwn", "*Details:*\n"+detail)
		if ind == 0 {
			detailsField := generateFields("mrkdwn", "*Details:*\n")
			detailsBlock := generateSectionBlock([]*Fields{detailsField})
			detailsBlocks = append(detailsBlocks, detailsBlock)
		}
		detailsText := generateText("plain_text", detail, &emoji)
		detailsBlock := generateSingleBlock("section",detailsText)
		detailsBlocks = append(detailsBlocks, detailsBlock)
	}

	//detailsField := generateFields("mrkdwn", "*Details:*\n"+detailsArr[0])
	//detailsBlock := generateSectionBlock([]*Fields{detailsField})
	//detailsBlocks = append(detailsBlocks, detailsBlock)
	//
	//detailsField = generateFields("mrkdwn", "*Details:*\n"+detailsArr[1])
	//detailsBlock = generateSectionBlock([]*Fields{detailsField})
	//detailsBlocks = append(detailsBlocks, detailsBlock)

	blocks := []Blocks{headerBlock, serviceInfoBlock, summaryBlock}

	for ind, block := range detailsBlocks {
		if ind < 30 {
			blocks = append(blocks, block)
		}
	}

	attachment := Attachments{
		Color: "#f2c744",
		Blocks: blocks,
	}

	slackBody, _ := json.Marshal(SlackRequestBody{Attachments: []Attachments{attachment}})
	req, err := http.NewRequest(http.MethodPost, sc.webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	if buf.String() != "ok" {
		fmt.Print(resp.Status, "\n")
		return errors.New("non-ok response returned from slack")
	}
	return nil
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}