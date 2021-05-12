package main

import (
	tg "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func getJidFromMsgText(s string) string {
	if s == "" {
		return ""
	}

	msgArray := strings.Split(s, "\n")
	if len(msgArray) < 1 {
		return ""
	}

	jid := strings.Trim(msgArray[0], " ")
	if !strings.HasPrefix(jid, "JID: ") {
		return ""
	}

	return jid[5:]
}

func getJidFromMessage(m *tg.Message) string {
	jid := getJidFromMsgText(m.ReplyTo.Text)

	if jid == "" {
		return getJidFromMsgText(m.ReplyTo.Caption)
	}
	return jid
}

func sendTelegramTxt(str string) error {
	_, err := bot.Send(
		tg.ChatID(chatId),
		str,
		tg.NoPreview,
		"Markdown")
	return err
}

func isMsgNeedProcess(m *tg.Message) bool {
	if m.Unixtime < startTime {
		return false
	}

	if m.Chat.ID != chatId {
		return false
	}

	if !m.IsReply() {
		_, _ = bot.Send(
			tg.ChatID(chatId),
			"Only accept reply.",
			tg.NoPreview,
			"Markdown")
		return false
	}
	return true
}
