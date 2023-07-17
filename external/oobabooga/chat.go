package oobabooga

import (
	"encoding/json"
	"fmt"
	"ss/network"
)

func GetAnswerFor(history []string, bio string) string {
	start := "you are pretending to be a human. Their biography is " + bio +
		` Using your biography, make a short reply as the human you are pretnending to be, speak in the first person as them.`

	items := []string{history[0]}
	fmt.Println(items)
	internal := []any{items}
	visible := []any{items}
	historyMap := map[string]any{"internal": internal, "visible": visible}

	m := map[string]any{"model": "ggml-vic7b-uncensored-q5_0.bin",
		"_continue":  false,
		"history":    historyMap,
		"user_input": start}

	fmt.Println(m)
	jsonString, code := network.PostTo("http://127.0.0.1:5000/api/v1/chat",
		"", m)
	if code != 200 {
		fmt.Println(jsonString, code)
		return ""
	}
	return parseJson(jsonString)
}

func SlackMessageFromBio(bio string) string {
	payload := makePayloadForBio(bio)
	jsonString, code := network.PostTo("http://127.0.0.1:5000/api/v1/chat",
		"", payload)
	if code != 200 {
		fmt.Println(jsonString, code)
		return ""
	}
	return parseJson(jsonString)
}

func makePayloadForBio(bio string) map[string]any {
	start := "you are pretending to be a human. Their biography is " + bio
	q := start + ". " + "make a slack message from the human you are pretending to be and use something from their background to make it personal and sound like it's coming from them. Don't use hashtags like #this. Don't make it over the top silly. Make a short thing about something that happened in their lives from a few years ago that's slightly amusing."
	history := map[string]any{"internal": []string{}, "visible": []string{}}
	m := map[string]any{"model": "ggml-vic7b-uncensored-q5_0.bin", "_continue": false,
		"history": history, "user_input": q}
	return m
}

func parseJson(jsonString string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r, jsonString)
		}
	}()
	var m map[string]any
	//fmt.Println(jsonString)
	json.Unmarshal([]byte(jsonString), &m)
	choices := m["results"].([]any)
	choice := choices[0].(map[string]any)
	history := choice["history"].(map[string]any)
	visible := history["visible"].([]any)
	list := visible[0].([]any)
	return list[len(list)-1].(string)
}
