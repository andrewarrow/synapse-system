package websocket

import (
	"encoding/json"
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
		log.Println("Received message:", string(message))
	}
}

func (s *SystemWs) SendSlack(token, user, channel string) {
	authMessage := map[string]string{}
	authMessage["type"] = "auth"
	authMessage["app_id"] = s.AppId
	authMessage["user_id"] = user
	authMessage["token"] = token
	asBytes, _ := json.Marshal(authMessage)

	err := s.Conn.WriteMessage(websocket.TextMessage, asBytes)
	if err != nil {
		log.Println("Failed to send message:", err)
	}

	/*
		message = {
		    'type': 'message',
		    'channel': 'channel_id',
		    'user': 'user_id',
		    'text': 'Hello, Slack!'
		}
	*/
}
