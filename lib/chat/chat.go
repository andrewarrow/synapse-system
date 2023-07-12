package chat

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

type Friend struct {
	Path     string
	FullName string
}

func NewFriend(user map[string]any) *Friend {
	f := Friend{}
	f.FullName = user["full_name"].(string)
	f.Path = firstNameLower(f.FullName)
	//fmt.Println(fullName, path)
	//token := user["slack_token"].(string)
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
		time.Sleep(time.Second * time.Duration(intSeconds))
	}
}

func Run(r *router.Router) {
	users := r.All("user", "order by created_at", "")
	for _, user := range users {
		friend := NewFriend(user)
		go friend.OnlineOffline()
		intSeconds := rand.Intn(10) + 6
		time.Sleep(time.Second * time.Duration(intSeconds))
	}
}

func turnGreenLightOn(path, seconds string) {
	b, err := exec.Command("python3", "python/online.py", "python/"+path, seconds).CombinedOutput()
	fmt.Println(string(b), err)
}

func firstNameLower(s string) string {
	tokens := strings.Split(s, " ")
	return strings.ToLower(tokens[0])
}
