package slack

import (
	"fmt"
	"ss/network"
)

func PostMessage(token, id, text string) {
	payload := map[string]any{"channel": id, "text": text}
	jsonString, code := network.PostTo("https://slack.com/api/chat.postMessage",
		token, payload)
	if code != 200 {
		fmt.Println(jsonString, code)
	}
}
func GetHistory(id string) {
	payload := map[string]any{"channel": id, "limit": 30}
	jsonString, code := network.PostTo("https://slack.com/api/conversations.history",
		apiKey, payload)
	if code != 200 {
		fmt.Println(jsonString, code)
	}
	fmt.Println(jsonString, code)
}
