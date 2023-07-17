package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"ss/external/oobabooga"
	"ss/external/slack"

	"github.com/andrewarrow/feedback/router"
	"github.com/gorilla/websocket"
)

type SystemWs struct {
	Conn  *websocket.Conn
	AppId string
}

func Connect(socketURL, appId string, c *router.Context) *SystemWs {
	s := SystemWs{}
	var err error
	s.Conn, _, err = websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatal("WebSocket connection error:", err)
		return nil
	}
	s.AppId = appId

	go readMessages(s.Conn, c)

	return &s
}

func readMessages(conn *websocket.Conn, c *router.Context) {
	users := c.All("user", "order by id", "")
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			return
		}
		jsonString := string(message)
		event, eventId := parseSlackSocketJson(jsonString)
		if event == nil {
			continue
		}
		c.Params = map[string]any{}
		c.Params["id_event"] = eventId
		msg := c.Insert("slack_event")
		if msg == "" {
			eventFlavor := event["type"].(string)
			if eventFlavor == "message" {
				user := event["user"].(string)
				channel := event["channel"].(string)
				go handleNewMessage(user, channel, users)
			}
		}
	}
}

func handleNewMessage(user, channel string, users []map[string]any) {
	if user != "U05G6HWQB7V" {
		return
	}
	jsonString := slack.GetHistory(channel)
	items := parseSlackHistoryJson(jsonString)

	theUser := users[rand.Intn(len(users))]

	a := oobabooga.GetAnswerFor(items, theUser["bio"].(string))
	slack.PostMessage(theUser["slack_token"].(string), channel, a)
}

func parseSlackHistoryJson(jsonString string) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	//fmt.Println(jsonString)

	buffer := []string{}
	messages := m["messages"].([]any)
	for _, thing := range messages {
		message := thing.(map[string]any)
		text := message["text"].(string)
		buffer = append(buffer, text)
	}
	return buffer
}

func parseSlackSocketJson(jsonString string) (map[string]any, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	//fmt.Println(jsonString)

	flavor := m["type"].(string)
	if flavor == "events_api" {
		payload := m["payload"].(map[string]any)
		event := payload["event"].(map[string]any)
		return event, payload["event_id"].(string)
	}
	return nil, ""
}
