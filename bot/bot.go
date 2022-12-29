package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var BotToken string

func Start() {

	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	//discord.AddHandler(Message)
	discord.Open()

	defer discord.Close()
	fmt.Print("Bot foi iniciado com sucesso")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
}
