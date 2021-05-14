package main

import (
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"time"
)

var (
	whatsConn         *whatsapp.Conn
	bot               *tg.Bot
	tgToken                 = ""
	skipNotifyMap           = make(map[string]bool)
	chatId            int64 = 0
	startTime               = time.Now().Unix()
	waShortClientName       = "KevinZonda's Workstation"
	waLongClientName        = "KevinZonda"
	waTimeout               = 20 * time.Second
)
