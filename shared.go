package main

import (
	tg "gopkg.in/tucnak/telebot.v2"
	"time"
)

var (
	bot, Kerr = tg.NewBot(
		tg.Settings{
			Token: "",
			Poller: &tg.LongPoller{
				Timeout: 5 * time.Second,
			},
		},
	)
	skipNotifyMap           = make(map[string]bool)
	chatId            int64 = 0
	startTime               = time.Now().Unix()
	waShortClientName       = "KevinZonda's Workstation"
	waLongClientName        = "KevinZonda"
	waTimeout               = 20 * time.Second
)
