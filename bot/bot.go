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

	case strings.Contains(message.Content, "!html"):
		discord.ChannelMessageSend(message.ChannelID,
			"https://developer.mozilla.org/pt-BR/docs/Web/HTML")

	case strings.Contains(message.Content, "!css"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do CSS: https://developer.mozilla.org/pt-BR/docs/Web/CSS")

	case strings.Contains(message.Content, "!js"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do JavaScript: https://developer.mozilla.org/pt-BR/docs/Web/JavaScript")
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é um dos melhores documentos explicativos do Javascript: https://javascript.info/")

	case strings.Contains(message.Content, "!javascript"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do JavaScript: https://developer.mozilla.org/pt-BR/docs/Web/JavaScript")
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é um dos melhores documentos explicativos do Javascript: https://javascript.info/")

	case strings.Contains(message.Content, "!node"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Node.js: https://nodejs.org/pt-br/docs/")
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Express.js: https://expressjs.com/pt-br/")

	case strings.Contains(message.Content, "!ts"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do TypeScript: https://www.typescriptlang.org/pt/docs/")

	case strings.Contains(message.Content, "!typescript"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do TypeScript: https://www.typescriptlang.org/pt/docs/")

	case strings.Contains(message.Content, "!react"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do React: https://react.dev/")

	case strings.Contains(message.Content, "!vite"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Vite: https://vitejs.dev/guide/")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Vite (npm) use: ```npm create vite@latest```")

	case strings.Contains(message.Content, "!Vite"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Vite: https://vitejs.dev/guide/")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Vite (npm) use: ```npm create vite@latest```")

	case strings.Contains(message.Content, "!next"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Next.js: https://nextjs.org/docs/getting-started")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Next.js (npm) use: ```npx create-next-app@latest```")

	case strings.Contains(message.Content, "!Next"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Next.js: https://nextjs.org/docs/getting-started")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Next.js (npm) use: ```npx create-next-app@latest```")

	case strings.Contains(message.Content, "!astro"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Astro: https://docs.astro.build/pt-br/getting-started/")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Astro (npm) use: ```npm create astro@latest```")

	case strings.Contains(message.Content, "!Astro"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do Astro: https://docs.astro.build/pt-br/getting-started/")
		discord.ChannelMessageSend(message.ChannelID,
			"Para iniciar um novo projeto com Astro (npm) use: ```npm create astro@latest```")

	case strings.Contains(message.Content, "!C#"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do C#: https://learn.microsoft.com/pt-br/dotnet/csharp/")
	case strings.Contains(message.Content, "!.NET"):
		discord.ChannelMessageSend(message.ChannelID,
			"Essa é a documentação do .NET: https://learn.microsoft.com/pt-br/dotnet/")
	}
}
