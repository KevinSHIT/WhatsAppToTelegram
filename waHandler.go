package main

import (
	"WhatsAppToTelegram/utils"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	tg "gopkg.in/tucnak/telebot.v2"
	"os"
	"path"
	"strings"
)

type waHandler struct {
	wac       *whatsapp.Conn
	startTime uint64
}

func (wh *waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error caught in handler: %v\n", err)
	if strings.Contains(err.Error(), "code: 1000") {
		errC := waConn.Restore()
		if errC != nil {
			fmt.Fprintf(os.Stderr, "error caught in handler: <Restore> %v\n", err)
		}
	}
}

func (wh *waHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	if message.Info.FromMe || message.Info.Timestamp < wh.startTime {
		return
	}

	jid := getJid(message.Info)

	isGroup := strings.Contains(jid, "-")

	if !utils.Exists("cache") {
		os.Mkdir("cache", os.ModePerm)
	}

	imgPath := path.Join("cache", utils.RandStringBytes(10)+".jpg")

	for {
		if !utils.Exists(imgPath) {
			break
		}
		imgPath = path.Join("cache", utils.RandStringBytes(10)+".jpg")
	}

	img, err := message.Download()

	if err != nil {
		if !isGroup {
			sendWhatsAppTxtMsg(
				wh,
				message.Info.RemoteJid,
				"Message transferred failed. More info please text .help\n"+
					"Err: WA_R_IMG_DL",
			)
		}
		_ = sendTelegramTxt(fmt.Sprintf(
			"JID: %s\n"+
				"MSG: IMG\n"+
				"ERR: Image download failed!", jid))
		return
	}

	err = utils.SavePic(img, imgPath)
	defer os.Remove(imgPath)

	if err != nil {
		if !isGroup {
			sendWhatsAppTxtMsg(
				wh,
				message.Info.RemoteJid,
				"Message transferred failed. More info please text .help\n"+
					"Err: WA_R_IMG_SV",
			)
		}

		_ = sendTelegramTxt(fmt.Sprintf(
			"JID: %s\n"+
				"MSG: IMG\n"+
				"ERR: Image save failed!", jid))
		return
	}

	_, err = tgBot.Send(
		tg.ChatID(tgChatId),
		&tg.Photo{
			File:    tg.FromDisk(imgPath),
			Caption: "JID: " + jid,
		},
	)

	if err != nil {
		if !isGroup {
			sendWhatsAppTxtMsg(
				wh,
				message.Info.RemoteJid,
				"Message transferred failed. More info please text .help\n"+
					"Err: TG_S_IMG",
			)
		}
		_ = sendTelegramTxt(fmt.Sprintf(
			"JID: %s\n"+
				"MSG: IMG\n"+
				"ERR: Image sent failed!", jid))
		fmt.Println(err)
		return
	}

	if !isGroup {
		if !skipNotify(message.Info.RemoteJid) {
			sendWhatsAppTxtMsg(
				wh,
				message.Info.RemoteJid,
				"Message transferred successfully. More info please text .help",
			)
		}
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

	if strings.Contains(jid, "-") {
		return
	}

	transferState := true

	errS := sendTelegramTxt(msgStr)

	if errS != nil {
		fmt.Fprintf(os.Stderr, "Cannot send to TG")
		transferState = false
	}

	// Func
	if strings.HasPrefix(message.Text, ".") {
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Text: getResponse(message.Text, message.Info.RemoteJid),
		}

		if _, err := wh.wac.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}

		return
	}

	if transferState && skipNotify(message.Info.RemoteJid) {
		return
	}

	transferStateMsg := "successfully"
	if !transferState {
		transferStateMsg = "failed"
	}

	sendWhatsAppTxtMsg(
		wh,
		message.Info.RemoteJid,
		fmt.Sprintf(
			"Message transferred %s. More info please text .help",
			transferStateMsg,
		),
	)
}

func skipNotify(id string) bool {
	v, exist := skipNotifyMap[id]
	if !exist {
		return false
	}
	return v
}
