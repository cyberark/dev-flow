package chat

import (
	"context"
	"os"
	
	"github.com/nlopes/slack"
)

type Slack struct{}

func (s Slack) client() *slack.Client {
	return slack.New(os.Getenv("SLACK_API_TOKEN"))
}

func (s Slack) getUserID(username string) (string) {
	users, err := s.getAllUsers()
	
	if err != nil {
		panic(err)
	}

	var userID string

	for _, slackUser := range users {
		if slackUser.Name == username || slackUser.Profile.DisplayName == username {
			userID = slackUser.ID
		}
	}
	
	return userID
}

func (s Slack) DirectMessage(username string, message string) {
	userID := s.getUserID(username)

	client := s.client()
	
	_, _, channelID, err := client.OpenIMChannel(userID)

	if err != nil {
		panic(err)
	}

	params := slack.PostMessageParameters{
		Username: "dev-flow",
		AsUser: true,
	}
	
	client.PostMessage(channelID, message, params)
}

func (s Slack) getAllUsers() (results []slack.User, err error) {
	// The Slack API  may require pagination in the future, in which case
	// this limit of 0 will no longer work.
	up := s.client().GetUsersPaginated(
		slack.GetUsersOptionPresence(false),
		slack.GetUsersOptionLimit(0),
	)

	up, err = up.Next(context.Background())

	results = append(results, up.Users...)
	
	return results, up.Failure(err)
}
