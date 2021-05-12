package main

import (
	tg "gopkg.in/tucnak/telebot.v2"
	"time"
)

var (
	bot, Kerr = tg.NewBot(
		tg.Settings{
			Token: "",
			Poller: &tg.LongPoller {
				Timeout: 5 * time.Second,
			},
		},
	)
	chatId int64 = 0
)
