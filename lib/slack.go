package lib

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
)

func PostToSlack(url, channel, message string) bool {
	payload := slack.Payload{
		Text:      message,
		Username:  "Notification",
		Channel:   channel,
		IconEmoji: ":wink:",
	}

	err := slack.Send(url, "", payload)
	if err != nil {
		fmt.Printf("Fail to send to slack, error(%s)", err)
		return false
	}

	return true
}
