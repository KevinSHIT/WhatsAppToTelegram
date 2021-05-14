package main

import (
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"time"
)

var (
	tgChatId      int64 = 0
	tgBot         *tg.Bot
	tgToken       = ""
	skipNotifyMap = make(map[string]bool)
	startTime     = time.Now().Unix()

	waConn            *whatsapp.Conn
	waShortClientName = "KevinZonda's Workstation"
	waLongClientName  = "KevinZonda"
	waTimeout         = 20 * time.Second
)
