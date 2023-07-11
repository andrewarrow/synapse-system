package chat

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func Run(r *router.Router) {
	users := r.All("user", "order by created_at", "")
	for _, user := range users {
		fullName := user["full_name"].(string)
		path := firstNameLower(fullName)
		fmt.Println(fullName, path)
		go turnGreenLightOn(path)
		token := user["slack_token"].(string)
		slackUser := user["slack_user"].(string)
		_, _ = token, slackUser
		//slack.SetPresence(token, slackUser)
		//slack.RtmConnect(token, slackUser)
		//slack.PostMessage(token, "C05G01UFYMU", "test")
	}
}

func turnGreenLightOn(path string) {
	b, err := exec.Command("python3", "python/online.py", "python/"+path, "26").CombinedOutput()
	fmt.Println(string(b), err)
}

func firstNameLower(s string) string {
	tokens := strings.Split(s, " ")
	return strings.ToLower(tokens[0])
}
