package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetReportCaller(true)
	for _, msg := range []string{":fearful: bad",":thumbsup: ok", ":crescent_moon: Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>"} {
		if err := MessageToSlack(msg); err != nil {
			logrus.Errorf("%+v", err)
		} else {
			logrus.Infof("%+v is sent", msg)
		}
	}

}

func MessageToSlack(msg string) error {
	env := "WEB_HOOK_URL_SLACK"
	webhookUrl := os.Getenv(env)
	if len(webhookUrl) == 0 {
		return errors.Errorf("no env var %+v", env)
	}
	payload := slack.Payload{
		Text:        msg,
		Username:    "robot",
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		return errors.Errorf("errors in sending to slack: %+v", err)
	}
	return nil
}