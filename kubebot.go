package main

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"regexp"
	"strings"
)

const (
	pattern = "(?i)\\b(k|kubectl)\\b"
)

var (
	re = regexp.MustCompile(pattern)
)

func kubernetes(command *bot.PassiveCmd) (string, error) {
	if !re.MatchString(command.Raw) {
		return "", nil
	}

	kubecmd := re.ReplaceAllString(command.Raw, "")
	kubecmd = strings.Trim(kubecmd, " ")
	params := strings.Split(kubecmd, " ")
	output := execute("kubectl", params...)

	return fmt.Sprintf("```\n%s\n```", output), nil
}

func init() {
	bot.RegisterPassiveCommand(
		"kubernetes",
		kubernetes)
}
