package main

import (
	"github.com/go-chat-bot/bot/slack"
	"os"
)

func main() {
	slack.Run(os.Getenv("KUBEBOT_SLACK_TOKEN"))
}
