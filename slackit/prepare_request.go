package slackit

import (
	"bitbucket.org/shadowchef/utils/methods"
	"time"
)

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

func PrepareAttachmentBody(req ClientRequest) []Attachments {

	serviceName := req.ServiceName
	summary := req.Summary
	metadata := req.Metadata
	details := req.Details
	status := req.Status

	color := StatusMap[Warning]

	if v, ok := StatusMap[status]; ok {
		color = v
	}

	currentTime := time.Now()

	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")

	emoji := true

	headerText := generateText("plain_text", "New Alert", &emoji)

	headerBlock := generateSingleBlock("header", headerText)

	serviceNameField := generateFields("mrkdwn", "*Service:*\n"+serviceName)

	serviceLogTimeField := generateFields("mrkdwn", "*Created At:*\n"+currentTimeStr)

	serviceInfoBlock := generateSectionBlock([]*Fields{serviceNameField, serviceLogTimeField})

	summaryField := generateFields("mrkdwn", "*Summary:*\n"+summary)

	summaryBlock := generateSectionBlock([]*Fields{summaryField})

	var detailsBlocks []Blocks
	detailsArr := methods.Chunks(details, 2000)

	for ind, detail := range detailsArr{
		if ind == 0 {
			detailsField := generateFields("mrkdwn", "*Details:*\n")
			detailsBlock := generateSectionBlock([]*Fields{detailsField})
			detailsBlocks = append(detailsBlocks, detailsBlock)
		}
		detailsText := generateText("mrkdwn", "```"+detail+"```", nil)
		detailsBlock := generateSingleBlock("section",detailsText)
		detailsBlocks = append(detailsBlocks, detailsBlock)
	}

	metadataText := generateText("mrkdwn","*Metadata:*\n"+ metadata, nil)
	metadataBlock := generateSingleBlock("section",metadataText)

	blocks := []Blocks{headerBlock, serviceInfoBlock, summaryBlock, metadataBlock}

	for ind, block := range detailsBlocks {
		if ind < 45 {
			blocks = append(blocks, block)
		}
	}

	attachment := Attachments{
		Color: color,
		Blocks: blocks,
	}

	return []Attachments{attachment}
}
