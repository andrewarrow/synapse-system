package chat

import (
	"fmt"
	"math/rand"
	"ss/external/openai"
	"ss/external/slack"
	"time"
)

type Friend struct {
	Path     string
	FullName string
	Token    string
	Bio      string
}

func NewFriend(user map[string]any) *Friend {
	f := Friend{}
	f.FullName = user["full_name"].(string)
	f.Path = firstNameLower(f.FullName)
	f.Bio = user["bio"].(string)
	//fmt.Println(fullName, path)
	f.Token = user["slack_token"].(string)
	//slackUser := user["slack_user"].(string)
	//_, _ = token, slackUser
	//slack.SetPresence(token, slackUser)
	//slack.RtmConnect(token, slackUser)
	//slack.PostMessage(token, "C05G01UFYMU", "test")
	return &f
}

func (f *Friend) OnlineOffline() {
	for {
		intSeconds := rand.Intn(43200) + 900
		seconds := fmt.Sprintf("%d", intSeconds)
		fmt.Println(f.Path, seconds)
		go turnGreenLightOn(f.Path, seconds)
		go f.DoStuffWhileOnline(time.Now().Unix() + int64(intSeconds))
		time.Sleep(time.Second * time.Duration(intSeconds))
		intSeconds = rand.Intn(21600) + 900
		time.Sleep(time.Second * time.Duration(intSeconds))
	}
}

func (f *Friend) DoStuffWhileOnline(endAt int64) {
	intSeconds := rand.Intn(10) + 6
	time.Sleep(time.Second * time.Duration(intSeconds))
	for {
		if rand.Intn(100) == 33 {
			go f.PostInGeneral()
		}
		time.Sleep(time.Second)
		if time.Now().Unix() > endAt {
			break
		}
	}
}

func (f *Friend) PostInGeneral() {
	txt := openai.SlackMessageFromBio(f.Bio)
	slack.PostMessage(f.Token, "C05G01UFYMU", txt)
}
