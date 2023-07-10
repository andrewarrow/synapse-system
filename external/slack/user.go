package slack

import (
	"fmt"
	"ss/network"
)

/*

-d '{
    "presence": "auto",
    "user": "'"$USER_ID"'"
  }' \
  https://slack.com/api/users.setPresence

*/

func SetPresence(token, id string) {
	payload := map[string]any{"presence": "auto", "user": id}
	jsonString, code := network.PostTo("https://slack.com/api/users.setPresence",
		token, payload)
	fmt.Println(jsonString, code)
}

// https://slack.com/api/rtm.connect
func RtmConnect(token, id string) {
	payload := map[string]any{"presence": "auto", "user": id}
	jsonString, code := network.PostTo("https://slack.com/api/rtm.connect",
		token, payload)
	fmt.Println(jsonString, code)
}
