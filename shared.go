package main

import (
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"time"
)

const (
	tgChatId          int64 = 0
	tgToken                 = ""
	waShortClientName       = "KevinZonda's Workstation"
	waLongClientName        = "KevinZonda"
	waTimeout               = 20 * time.Second
)

var (
	tgBot         *tg.Bot
	skipNotifyMap = make(map[string]bool)
	startTime     = time.Now().Unix()
	waConn        *whatsapp.Conn
)
