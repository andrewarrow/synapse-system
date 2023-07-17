package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type SystemWs struct {
	Conn  *websocket.Conn
	AppId string
}

func Connect(socketURL, appId string) *SystemWs {
	s := SystemWs{}
	var err error
	s.Conn, _, err = websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatal("WebSocket connection error:", err)
		return nil
	}
	s.AppId = appId

	go readMessages(s.Conn)

	return &s
}

func readMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			return
		}
		jsonString := string(message)
		user, channel := parseSlackSocketJson(jsonString)
		go handleNewMessage(user, channel)
	}
}

func handleNewMessage(user, channel string) {
}

func parseSlackSocketJson(jsonString string) (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)

	flavor := m["type"].(string)
	if flavor == "events_api" {
		payload := m["payload"].(map[string]any)
		event := payload["event"].(map[string]any)
		user := event["user"].(string)
		channel := event["channel"].(string)
		return user, channel
	}
	return "", ""
}
