package main

import (
	tg "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func getJidFromMsgText(s string) string {
	if s == "" {
		return ""
	}

	if !strings.HasPrefix(s, "JID: ") {
		return ""
	}

	msgArray := strings.Split(s, "\n")
	if len(msgArray) < 1 {
		return ""
	}

	return strings.Trim(msgArray[0][5:], " ")
}

func sendTelegramTxt(str string) error {
	_, err := bot.Send(
		tg.ChatID(chatId),
		str,
		tg.NoPreview,
		"Markdown")
	return err
}
