package main

import "strings"

func getJid(s string) string {
	if s == "" {
		return ""
	}

	msgArray := strings.Split(s, "\n")
	if len(msgArray) < 1 {
		return ""
	}

	jid := strings.Trim(msgArray[0], " ")
	if !strings.HasPrefix(jid, "JID: ") {
		return ""
	}

	return  s[5:]
}
