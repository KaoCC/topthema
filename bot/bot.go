package bot

import (
	"bytes"
	"log"
	"text/template"
	"topthema/util"

	"github.com/slack-go/slack"
)

type bot struct {
	client *slack.Client
	tmpl   *template.Template
}

func New() *bot {
	authToken := util.ReadEnv("OAUTH_TOKEN")
	api := slack.New(authToken)

	tmpl := template.New("URL")
	tmpl, err := tmpl.Parse("<{{.Link}}|{{.Title}}>")
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	return &bot{
		client: api,
		tmpl:   tmpl,
	}
}

func (bot *bot) Post(payload util.Record) {

	CHANNEL_ID := "general"

	var buff bytes.Buffer
	if err := bot.tmpl.Execute(&buff, payload); err != nil {
		log.Fatal("Execute: ", err)
	}

	attachment := slack.Attachment{
		Text: buff.String(),
	}

	channelId, timestamp, err := bot.client.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText(payload.Title, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)

}
