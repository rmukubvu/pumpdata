package notifications

import "github.com/bluele/slack"

const (
	token       = ""
	channelName = "amakosi-ops"
)

func SlackSend(message string) error {
	api := slack.New(token)
	err := api.ChatPostMessage(channelName, message, nil)
	if err != nil {
		return err
	}
	return nil
}
