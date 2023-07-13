package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"ss/external/openai"
	"ss/external/slack"
	"ss/lib/chat"
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
		chat.Run(r)
	} else if arg == "test" {
		//slack.PostMessage("C05G01UFYMU", "test")
		txt := openai.SlackMessageFromBio(`Fred Clark was born in Boston, Massachusetts in 1979. He left there when he was 14 but still speaks with a Boston accent. His dad was a military man and moved around a lot, divorced from his mom Fred lived with her most of the time. He was a troubled youth, skipping school, doing drugs, but eventually at age 19 found his way with computer programming. As a high school dropout he had some trouble at first getting jobs, but his skills as a coder made him desirable to hire. By age 24 he was leading teams at top silicon valley tech startups. Today he works for many.pw as a software engineer writing mainly in the language go and python. He has a wicked smart sense of humor and loves to post funny things in slack when the time is right. He also works very hard and does not post too much.`)
		fmt.Println(txt)

	}

	select {}
}
