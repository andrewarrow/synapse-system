package slack

import (
	"encoding/json"
	"fmt"
	"ss/network"
	"strings"
)

// https://slack.com/api/apps.connections.open

func ConnectionsOpen() (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	payload := map[string]any{"token": appLevel}
	jsonString, code := network.PostTo("https://slack.com/api/apps.connections.open",
		appLevel, payload)
	fmt.Println(jsonString, code)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	url := m["url"].(string)
	tokens := strings.Split(url, "app_id=")
	appId := tokens[len(tokens)-1]

	return url, appId
}
