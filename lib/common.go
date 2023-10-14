package lib

import (
	"fmt"
	"os"
)

var SlackWebhook = os.Getenv("SLACK_WEBHOOK")
var SlackChannel = os.Getenv("SLACK_CHANNEL")

func HandleErr(err error) {
	if err != nil {
		errorMsg := fmt.Sprintf("Error : %s", err.Error())
		fmt.Println(errorMsg)
		PostToSlack(SlackWebhook, SlackChannel, errorMsg)
	}
}
