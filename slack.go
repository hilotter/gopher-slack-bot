package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Slack struct {
	WebhookUrl string
	Payload    Payload
}

type Payload struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	IconUrl   string `json:"icon_url"`
}

func (s *Slack) Post() (res *http.Response, err error) {
	params, err := json.Marshal(s.Payload)
	if err != nil {
		return nil, err
	}

	vs := url.Values{}
	vs.Add("payload", string(params))

	res, errp := http.PostForm(s.WebhookUrl, vs)
	if errp != nil {
		return nil, errp
	}

	return res, nil
}
