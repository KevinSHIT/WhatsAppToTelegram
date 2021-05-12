package main

import (
	"encoding/gob"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
)

func login(wac *whatsapp.Conn) error {
	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring session failed: %v", err)
		}
	} else {
		qr := make(chan string)

		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()

		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	if err = writeSession(session); err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}

	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(&session); err != nil {
		return session, err
	}

	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(session); err != nil {
		return err
	}
	return nil
}
