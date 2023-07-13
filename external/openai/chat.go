package openai

import (
	"encoding/json"
	"fmt"
	"ss/network"
)

func SlackMessageFromBio(bio string) string {
	payload := makePayloadForBio(bio)
	jsonString, code := network.PostTo("https://api.openai.com/v1/chat/completions",
		apiKey, payload)
	if code != 200 {
		fmt.Println(jsonString, code)
		return ""
	}
	return parseOpenAiJson(jsonString)
}

func makePayloadForBio(bio string) map[string]any {
	return nil
}

func parseOpenAiJson(jsonString string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var m map[string]any
	//fmt.Println(jsonString)
	json.Unmarshal([]byte(jsonString), &m)
	choices := m["choices"].([]any)
	choice := choices[0].(map[string]any)
	message := choice["message"].(map[string]any)
	c := message["content"].(string)
	return c
}
