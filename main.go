package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
)

func main() {
	if Kerr != nil {
		fmt.Fprintf(os.Stderr, "error creating tg connection: %v\n", Kerr)
		return
	}

	_, _ = bot.Send(
		tg.ChatID(chatId),
		fmt.Sprintf("Bot is connected"),
		tg.NoPreview,
		"Markdown",
	)

	wac, err := whatsapp.NewConnWithOptions(&whatsapp.Options{
		Timeout:         waTimeout,
		ShortClientName: waShortClientName,
		LongClientName:  waLongClientName,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	wac.AddHandler(&waHandler{
		wac,
		uint64(time.Now().Unix()),
	})

	if err = login(wac); err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	/* FIXED: seems not work */
	bot.Handle(tg.OnText, func(m *tg.Message) {
		if m.Unixtime < startTime {
			return
		}

		if m.Chat.ID != chatId {
			return
		}

		if !m.IsReply() {
			_, _ = bot.Send(
				tg.ChatID(chatId),
				"Only accept reply.",
				tg.NoPreview,
				"Markdown")
			return
		}

		jid := getJidFromMsgText(m.ReplyTo.Text)

		if jid == "" {
			jid = getJidFromMsgText(m.ReplyTo.Caption)
			if jid == "" {
				return
			}
		}

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: jid,
			},
			Text: m.Text,
		}

		if _, err := wac.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}
	})

	go bot.Start()
	<-time.After(360 * 24 * time.Hour)
}
