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
	messages := makeListOfMessages(bio)
	m := map[string]any{"model": "gpt-4",
		"messages": messages}
	return m
}

func makeListOfMessages(bio string) []any {
	system := map[string]any{"content": "you are pretending to be a human. Their biography is " + bio, "role": "system"}
	user := map[string]any{"content": "make a work related slack message from the human you are pretending to be", "role": "user"}

	list := []any{system, user}
	return list
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
