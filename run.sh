#!/bin/bash

binary="./WhatsAppToTelegram"
if [ ! -d "$binary" ];then
  ./build.sh
fi
chmod +x "$binary"
$binary