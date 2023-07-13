package openai

import "os"

var apiKey string

func Init() {
	apiKey = os.Getenv("OPEN_AI")
}
