package chat

import (
	"os"
)

type ChatClient interface {
	DirectMessage(string, string)
}

func GetClient() ChatClient {
	if os.Getenv("SLACK_API_TOKEN") != "" {
		return Slack{}
	} else {
		return nil
	}
}
