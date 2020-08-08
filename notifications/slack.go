package notifications

import "github.com/bluele/slack"

const (
	token       = "xoxb-1215817235969-1200780626695-Me0HuIvp450n52dn1UcFdwEk"
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
