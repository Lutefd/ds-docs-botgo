package main

import (
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
