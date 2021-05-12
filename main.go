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

	wac, err := whatsapp.NewConn(20 * time.Second)
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

	/* FIXME: seems not work */
	bot.Handle(tg.OnText, func(m *tg.Message) {
		fmt.Println(m.Text)
		if m.Chat.ID != 573387497 {
			return
		}
		println(m.IsReply())
		if !m.IsReply() {
			_, _ = bot.Send(
				tg.ChatID(chatId),
				"Only accept reply.",
				tg.NoPreview,
				"Markdown")
			return
		}

		jid := getJidFromMsgText(m.ReplyTo.Text)
		fmt.Println(jid)
		if jid == "" {
			return
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

	<-time.After(360 * 24 * time.Hour)
}
