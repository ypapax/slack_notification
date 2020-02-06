package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetReportCaller(true)
	if err := func() error {
		env := "WEB_HOOK_URL_SLACK"
		webhookUrl := os.Getenv(env)
		if len(webhookUrl) == 0 {
			return errors.Errorf("no env var %+v", env)
		}
		attachment1 := slack.Attachment{}
		attachment1.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
		attachment1.AddAction(slack.Action{Type: "button", Text: "Book flights 🛫", Url: "https://flights.example.com/book/r123456", Style: "primary"})
		attachment1.AddAction(slack.Action{Type: "button", Text: "Cancel", Url: "https://flights.example.com/abandon/r123456", Style: "danger"})
		payload := slack.Payload{
			Text:        "Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>",
			Username:    "robot",
			Channel:     "#general",
			IconEmoji:   ":monkey_face:",
			Attachments: []slack.Attachment{attachment1},
		}
		err := slack.Send(webhookUrl, "", payload)
		if len(err) > 0 {
			return errors.Errorf("errors in sending to slack: %+v", err)
		}
		return nil
	}(); err != nil {
		logrus.Fatalf("%+v", err)
	}
	logrus.Infof("sent")
}
