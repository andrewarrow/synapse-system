package main

import (
	_ "embed"
	"math/rand"
	"os"
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

	r := router.NewRouter("DATABASE_URL", embeddedFile)
	if arg == "init" {
	} else if arg == "run" {
		chat.Run(r)
	} else if arg == "test" {
		//slack.PostMessage("C05G01UFYMU", "test")
	}

	select {}
}
