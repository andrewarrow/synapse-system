package main

import (
	_ "embed"
	"math/rand"
	"os"
	"ss/external/slack"
	"ss/lib/cookies"
	"time"

	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

func main() {

	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	slack.Init()

	r := router.NewRouter("DATABASE_URL", embeddedFile)
	if arg == "init" {
	} else if arg == "run" {
		users := r.All("user", "order by created_at", "")
		for _, user := range users {
			token := user["slack_token"].(string)
			slackUser := user["slack_user"].(string)
			_, _ = token, slackUser
			//slack.SetPresence(token, slackUser)
			//slack.RtmConnect(token, slackUser)
			//slack.PostMessage(token, "C05G01UFYMU", "test")
		}
	} else if arg == "test" {
		//slack.PostMessage("C05G01UFYMU", "test")
		//email := os.Args[2]
		//chrome.Run(email)
		cookies.Parse("ttt")
	}

	select {}
}
