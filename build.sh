#!/bin/bash

rm -f shared.go
wget https://github.com/KevinZonda/WhatsAppToTelegram/raw/master/shared.go
nano shared.go
go build ./