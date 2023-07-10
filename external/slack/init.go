package slack

import "os"

var apiKey string
var appLevel string

func Init() {
	apiKey = os.Getenv("SLACK")
	appLevel = os.Getenv("SLACK_APP")
}
