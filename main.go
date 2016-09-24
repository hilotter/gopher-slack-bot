package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var (
		webhookUrl string
		username   string
		channel    string
		iconUrl    string
		iconEmoji  string
		text       string
	)

	envUrl := os.Getenv("SLACK_WEBHOOK_URL")
	flag.StringVar(&webhookUrl, "url", envUrl, "[required] slack webhook url")
	flag.StringVar(&webhookUrl, "u", envUrl, "[required] slack webhook url (short)")
	flag.StringVar(&username, "name", "gopher", "username")
	flag.StringVar(&username, "n", "gopher", "username (short)")
	flag.StringVar(&channel, "channel", "general", "channel")
	flag.StringVar(&channel, "c", "general", "channel (short)")
	flag.StringVar(&iconUrl, "iconurl", "", "icon url")
	flag.StringVar(&channel, "iu", "", "icon url (short)")
	flag.StringVar(&iconUrl, "iconemoji", "", "icon emoji")
	flag.StringVar(&channel, "ie", "", "icon emoji (short)")
	flag.StringVar(&text, "text", "", "[required] text")
	flag.StringVar(&text, "t", "", "[required] text (short)")
	flag.Parse()

	if webhookUrl == "" {
		flag.Usage()
		log.Fatal("url required or set SLACK_WEBHOOK_URL env")
	}
	if text == "" {
		flag.Usage()
		log.Fatal("text required")
	}
	s := &Slack{
		WebhookUrl: webhookUrl,
		Payload: Payload{
			Channel:   channel,
			Username:  username,
			Text:      text,
			IconUrl:   iconUrl,
			IconEmoji: iconEmoji,
		},
	}
	_, err := s.Post()
	if err != nil {
		log.Fatal(err)
	}
}
