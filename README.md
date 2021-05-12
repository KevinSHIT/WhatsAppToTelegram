# WhatsAppToTelegram

## How to use

### 1. Setup Environment

Before all, you should install the Go. **Go 1.16.3+** is recommended (cuz I only tested with it).

### 2. Modify Secrets

Then modify the information inside the `shared.go` to yours.

- `Token: ""`  
Modify to your Telegram bot token. If you don't have one, ask [@BotFather](https://t.me/BotFather).
- `chatId = 0`  
Modify to your Telegram ChatID. If you are not sure about your chatId, use [@userinfobot](https://t.me/userinfobot).
The first line starts with `Id` will be your chatId.

### 3. Build & Run

Use following commands to compile and run:

```bash
go build ./
chmod +x WhatsAppToTelegram
./WhatsAppToTelegram
```

## Functions Work

-[ ] WhatsApp -> Telegram
  -[x] Text Message
  -[x] Image Message
  -[ ] Document Message
  -[ ] Video Message
  -[ ] Audio Message
  -[ ] ~~Json Message~~
  -[ ] Contact Message
  -[ ] ~~Battery Message~~
  -[ ] New Contact
-[ ] Telegram -> WhatsApp
  -[x] Text Message
  -[ ] Photo Message
  -[ ] Document Message
  -[ ] ~~Video Message~~
  -[ ] Voice Message

## Credits

Thank you to all the repositories listed below (alphabetical ordered).

| Repository | License |
| --- | :---:|
| <https://github.com/Baozisoftware/qrcode-terminal-go> | BSD-3-Clause License |
| <https://github.com/Rhymen/go-whatsapp> | MIT |
| <https://gopkg.in/tucnak/telebot.v2> | MIT |
