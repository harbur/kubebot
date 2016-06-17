package main

import (
	"github.com/go-chat-bot/bot/slack"
	"os"
)

const (
	slackTokenLabel    string = "KUBEBOT_SLACK_TOKEN"
	slackChannelsLabel string = "KUBEBOT_SLACK_CHANNELS_IDS"
	slackAdminsLabel   string = "KUBEBOT_SLACK_ADMINS_NICKNAMES"
	slackCommandsLabel string = "KUBEBOT_SLACK_VALID_COMMANDS"
)

var (
	kb *Kubebot
)

func main() {
	kb = &Kubebot{
		token:    os.Getenv(slackTokenLabel),
		admins:   stringToMap(os.Getenv(slackAdminsLabel), " "),
		channels: stringToMap(os.Getenv(slackChannelsLabel), " "),
		commands: stringToMap(os.Getenv(slackCommandsLabel), " "),
	}

	slack.Run(kb.token)
}
