package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
)

var BotToken string

func Start() {

	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	discord.AddHandler(newMessage)

	discord.Open()

	defer discord.Close()
	fmt.Print("Bot foi iniciado com sucesso")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "docs"):
		discord.ChannelMessageSend(message.ChannelID,
			"Eu posso te ajudar com isso! Use !<nome da"+
				" tecnologia> para receber o link"+
				" da sua documentação")
	case strings.Contains(message.Content, "!react"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do React: https://beta.reactjs.org/")

	case strings.Contains(message.Content, "!css"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do CSS: https://developer.mozilla.org/pt-BR/docs/Web/CSS")

	}

}
