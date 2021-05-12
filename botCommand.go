package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getResponse(msg string, jid string) string {
	switch msg {
	case ".help":
		return "Hi. This is Zonda. The Transfer Bot!\n" +
			"Now the bot provides msg transferring function!\n" +
			"The source code will be released on my GitHub account under OKZPL or MIT asap.\n" +
			"If you are facing urgent situation, please send message through your real phone or call directly.\n" +
			"Sorry for the inconvenience caused by this. Have a good day! :)\n" +
			"One more thing, here is the command list:\n" +
			".help Get Help\n" +
			".setnotify [false|true] Set the notify which after transferring. If the forward is failed, the notify will still be sent."
	}
	if strings.HasPrefix(strings.ToLower(msg), ".setnotify") {
		b, err := strconv.ParseBool(strings.Trim(msg[10:], "x"))
		if err != nil {
			return "ERR: WA_P_TYPECOV"
		}
		if jid == "" {
			return "ERR: WA_P_JID_EMPTY"
		}
		skipNotifyMap[jid] = !b
		return fmt.Sprintf("Success set Notify to %v", b)
	}
	return "Sorry, cannot found the command!"

}
