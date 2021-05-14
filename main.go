package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var err error
	tgBot, err = tg.NewBot(
		tg.Settings{
			Token: tgToken,
			Poller: &tg.LongPoller{
				Timeout: 5 * time.Second,
			},
		},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating tg connection: %v\n", err)
		return
	}

	_, _ = tgBot.Send(
		tg.ChatID(tgChatId),
		fmt.Sprintf("Bot is connected"),
		tg.NoPreview,
		"Markdown",
	)

	waConn, err = whatsapp.NewConnWithOptions(&whatsapp.Options{
		Timeout:         waTimeout,
		ShortClientName: waShortClientName,
		LongClientName:  waLongClientName,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	waConn.AddHandler(&waHandler{
		waConn,
		uint64(time.Now().Unix()),
	})

	if err = login(waConn); err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	/* FIXED: seems not work */
	tgBot.Handle(tg.OnText, func(m *tg.Message) {
		if !isMsgNeedProcess(m) {
			return
		}

		jid := getJidFromMessage(m)

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: jid,
			},
			Text: m.Text,
		}

		if _, err := waConn.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}
	})

	tgBot.Handle(tg.OnPhoto, func(m *tg.Message) {
		if !isMsgNeedProcess(m) {
			return
		}

		imgPath := m.Photo.FilePath
		imgBytes := fileToBytes(imgPath)

		if imgBytes == nil {
			// TODO: Invalid Bytes
			return
		}

		jid := getJidFromMessage(m)

		msg := whatsapp.ImageMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: jid,
			},
			Thumbnail: imgBytes,
		}
		if _, err := waConn.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}

	})

	go tgBot.Start()
	<-time.After(360 * 24 * time.Hour)
}
