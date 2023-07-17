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

/*
{
  "envelope_id": "f9bdd0ed-09de-45f3-adb8-3f6144271f17",
  "payload": {
    "token": "I65hh46Ika7o3Z1xm3TYSSxh",
    "team_id": "T05G01UF8DC",
    "context_team_id": "T05G01UF8DC",
    "context_enterprise_id": null,
    "api_app_id": "A05GWDCG6GY",
    "event": {
      "client_msg_id": "66d1dc28-1e86-4429-9d27-9bee09dc0eec",
      "type": "message",
      "text": "general",
      "user": "U05G6HWQB7V",
      "ts": "1689602097.278579",
      "blocks": [
        {
          "type": "rich_text",
          "block_id": "B6U",
          "elements": [
            {
              "type": "rich_text_section",
              "elements": [
                {
                  "type": "text",
                  "text": "general"
                }
              ]
            }
          ]
        }
      ],
      "team": "T05G01UF8DC",
      "channel": "C05G01UFYMU",
      "event_ts": "1689602097.278579",
      "channel_type": "channel"
    },
    "type": "event_callback",
    "event_id": "Ev05HC82C42F",
    "event_time": 1689602097,
    "authorizations": [
      {
        "enterprise_id": null,
        "team_id": "T05G01UF8DC",
        "user_id": "U05GKBX505P",
        "is_bot": false,
        "is_enterprise_install": false
      }
    ],
    "is_ext_shared_channel": false,
    "event_context": "4-eyJldCI6Im1lc3NhZ2UiLCJ0aWQiOiJUMDVHMDFVRjhEQyIsImFpZCI6IkEwNUdXRENHNkdZIiwiY2lkIjoiQzA1RzAxVUZZTVUifQ"
  },
  "type": "events_api",
  "accepts_response_payload": false,
  "retry_attempt": 0,
  "retry_reason": ""
}
*/
