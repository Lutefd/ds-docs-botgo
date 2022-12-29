package main

import (
	"github.com/Lutefd/ds-docs-botgo/bot"
	"log"
	"os"
)

func main() {

	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("BOT_TOKEN não encontrado")
	}
	bot.BotToken = botToken
	bot.Start()
}
