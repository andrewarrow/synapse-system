package slack

import (
	"fmt"
	"ss/network"
)

/*
curl -X POST -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -H "Content-Type: application/json" -d '{
  "channel": "CHANNEL_ID",
  "text": "Your message"
}' https://slack.com/api/chat.postMessage

*/

func PostMessage(token, id, text string) {
	payload := map[string]any{"channel": id, "text": text}
	jsonString, code := network.PostTo("https://slack.com/api/chat.postMessage",
		token, payload)
	fmt.Println(jsonString, code)
}
