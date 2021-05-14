package main

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"os"
)

func tgOnText(m *tg.Message) {
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
}

func tgOnPhoto(m *tg.Message) {
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

}
