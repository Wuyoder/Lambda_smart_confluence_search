package lib

import (
	"fmt"
)

var SlackWebhook = ""
var SlackChannel = "#slack_web_hook_test"

func HandleErr(err error) {
	if err != nil {
		errorMsg := fmt.Sprintf("Error : %s", err.Error())
		fmt.Println(errorMsg)
		PostToSlack(SlackWebhook, SlackChannel, errorMsg)
	}
}
