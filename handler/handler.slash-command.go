package handler

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

func HandleSlashCommand(command slack.SlashCommand, client *slack.Client) error {

	switch command.Command {
	case "/hello":
		return handleHelloCommand(command, client)
	case "/was-this-article-useful":
		return handlerWasThisArticleUseful(command, client)
	}
	return nil
}

func handleHelloCommand(command slack.SlashCommand, client *slack.Client) error {

	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: command.UserName,
		},
	}

	attachment.Text = fmt.Sprintf("Hello %s", command.Text)
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed post message : %v", err)
	}

	return nil
}

func handlerWasThisArticleUseful(command slack.SlashCommand, client *slack.Client) error {

	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: command.UserName,
		},
	}

	attachment.Text = fmt.Sprintf("Hello %s", command.Text)
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed post message : %v", err)
	}

	return nil
}
