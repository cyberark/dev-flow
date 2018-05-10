package chat

type ChatClient interface {
	DirectMessage(string, string)
}

func GetClient() ChatClient {
	return Slack{}
}
