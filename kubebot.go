package main

import (
	"fmt"
	"time"

	"github.com/go-chat-bot/bot"
)

type Kubebot struct {
	token    string
	admins   map[string]bool
	channels map[string]bool
	commands map[string]bool
}

const (
	forbiddenUserMessage     string = "%s - ⚠ kubectl forbidden for user @%s\n"
	forbiddenChannelMessage  string = "%s - ⚠ Channel %s forbidden for user @%s\n"
	forbiddenCommandMessage  string = "%s - ⚠ Command %s forbidden for user @%s\n"
	forbiddenUserResponse    string = "Sorry @%s, but you don't have permission to run this command :confused:"
	forbiddenChannelResponse string = "Sorry @%s, but I'm not allowed to run this command here :zipper_mouth_face:"
	forbiddenCommandResponse string = "Sorry @%s, but I cannot run this command. I'm allowed to run `%s`"
	okResponse               string = "Roger that!\n@%s, this is the response to your request:\n ```\n%s\n``` "
)

func kubectl(command *bot.Cmd) (msg string, err error) {
	t := time.Now()
	time := t.Format(time.RFC3339)
	nickname := command.User.Nick

	if !kb.admins[nickname] {
		fmt.Printf(forbiddenUserMessage, time, nickname)
		return fmt.Sprintf(forbiddenUserResponse, nickname), nil
	}

	if !kb.channels[command.Channel] {
		fmt.Printf(forbiddenChannelMessage, time, command.Channel, nickname)
		return fmt.Sprintf(forbiddenChannelResponse, nickname), nil
	}

	if len(command.Args) > 0 && !kb.commands[command.Args[0]] {
		fmt.Printf(forbiddenCommandMessage, time, kb.commands, nickname)
		return fmt.Sprintf(forbiddenCommandResponse, nickname, kb.commands), nil
	}

	output := execute("kubectl", command.Args...)

	return fmt.Sprintf(okResponse, nickname, output), nil
}

func init() {
	bot.RegisterCommand(
		"kubectl",
		"Kubectl Slack integration",
		"",
		kubectl)
}
