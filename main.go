package main

import (
	log "github.com/kirillDanshin/dlog"
	tg "github.com/toby3d/telegram"
)

// bot is general structure of the bot
var bot *tg.Bot

// main function is a general function for work of this bot
func main() {
	log.Ln("Let'g Get It Started...")
	log.Ln("Initializing new bot via checking access_token...")
	var err error
	bot, err = tg.NewBot(cfg.UString("telegram.token"))
	errCheck(err)

	log.Ln("Let's check updates channel!")
	for update := range getUpdatesChannel() {
		switch {
		case update.IsInlineQuery():
			updateInlineQuery(update.InlineQuery)
		case update.IsMessage():
			updateMessage(update.Message)
		default:
			log.D(update)
		}
	}
}
