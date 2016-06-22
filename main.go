package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-chat-bot/bot/slack"
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

func validateEnvParams() error {
	if os.Getenv(slackTokenLabel) == "" {
		return errors.New("slackTokenLabel env variable not defined")
	}
	if os.Getenv(slackChannelsLabel) == "" {
		return errors.New("slackChannelsLabel env variable not defined")
	}
	if os.Getenv(slackAdminsLabel) == "" {
		return errors.New("slackAdminsLabel env variable not defined")
	}
	if os.Getenv(slackCommandsLabel) == "" {
		return errors.New("slackCommandsLabel env variable not defined")
	}

	return nil
}

func main() {

	if err := validateEnvParams(); err != nil {
		fmt.Printf("Kubebot cannot run. Error: %s\n", err.Error())
		return
	}

	kb = &Kubebot{
		token:    os.Getenv(slackTokenLabel),
		admins:   stringToMap(os.Getenv(slackAdminsLabel), " "),
		channels: stringToMap(os.Getenv(slackChannelsLabel), " "),
		commands: stringToMap(os.Getenv(slackCommandsLabel), " "),
	}

	slack.Run(kb.token)
}
