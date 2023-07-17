package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"ss/external/oobabooga"
	"ss/external/openai"
	"ss/external/slack"
	"ss/lib/chat"
	"ss/lib/websocket"
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
	openai.Init()

	r := router.NewRouter("DATABASE_URL", embeddedFile)
	if arg == "init" {
	} else if arg == "run" {
		wsUrl, appId := slack.ConnectionsOpen()
		fmt.Println(appId, wsUrl)
		websocket.Connect(wsUrl, appId, r.ToContext())
		chat.Run(r)
	} else if arg == "test" {
		//slack.PostMessage("C05G01UFYMU", "test")
		txt := oobabooga.SlackMessageFromBio(`Jenny Wey was born in London and lived there until the age of 24 when she moved to New York City with her fiancé mainly for his job. After just a year they broke up but she fell in love with the city and stayed there. Having an ocean between her and her family back home was actually nice for her. It created sound good boundaries and let her live her own life. She got into project management for tech companies and learned how to work with software developers very well. She does not know how to code anything herself but she understands the process and the terms and knows how to set direction and get the product from point A to point B as quickly as possible. Part of this is having fun on slack and just chatting with the team to make them feel relaxed and having fun at work. But there is a balance she is great at. She does not post too much in slack, just the right amount.`)
		fmt.Println(txt)

	}

	select {}
}
