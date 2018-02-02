package main

import tg "github.com/toby3d/telegram"

func updateMessage(message *tg.Message) {
	if message.IsCommand("") {
		commands(message)
		return
	}

	messages(message)
}
