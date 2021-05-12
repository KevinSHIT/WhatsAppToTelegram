package main

func getResponse(s string) string {
	switch s {
	case ".help":
		return "Hi. This is Zonda. The Transfer Bot!\n" +
			"Now the bot provides msg transferring function!\n" +
			"The source code will be released on my GitHub account under OKZPL or MIT asap.\n" +
			"If you are facing urgent situation, please send message through your real phone or call directly.\n" +
			"Sorry for the inconvenience caused by this. Have a good day! :)"
	}
	return "Sorry, cannot found the command!"

}
