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

	tgBot.Handle(tg.OnText, tgOnText)

	tgBot.Handle(tg.OnPhoto, tgOnPhoto)

	go tgBot.Start()
	<-time.After(360 * 24 * time.Hour)
}
