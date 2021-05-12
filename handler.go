package main

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"os"
	"strings"
)

type waHandler struct {
	wac       *whatsapp.Conn
	startTime uint64
}

func (wh *waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error caught in handler: %v\n", err)
}

func getJid(info whatsapp.MessageInfo) string {
	if info.Source.Participant == nil {
		return info.RemoteJid
	} else {
		return *info.Source.Participant
	}
}

func (wh *waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if message.Info.FromMe || message.Info.Timestamp < wh.startTime {
		return
	}

	jid := getJid(message.Info)

	msgStr := fmt.Sprintf(
		"JID: %s\nMsg: %s\n",
		jid,
		message.Text,
	)

	fmt.Fprintf(os.Stdout, msgStr)

	_, errS := bot.Send(
		tg.ChatID(chatId),
		msgStr,
		tg.NoPreview,
		"Markdown")

	transferState := "successfully"

	if errS != nil {
		fmt.Fprintf(os.Stderr, "Cannot send to TG")
		transferState = "failed"
	}

	if !strings.Contains(jid, "-") {
		if strings.HasPrefix(message.Text, ".") {
			msg := whatsapp.TextMessage{
				Info: whatsapp.MessageInfo{
					RemoteJid: message.Info.RemoteJid,
				},
				Text: getResponse(message.Text),
			}

			if _, err := wh.wac.Send(msg); err != nil {
				fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
			}

			return
		}

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Text: fmt.Sprintf(
				"Message transferred %s. More info please text .help",
				transferState,
			),
		}

		if _, err := wh.wac.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}
	}
}
