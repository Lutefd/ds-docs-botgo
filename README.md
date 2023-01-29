# Languages and Frameworks documentation bot

This bot was created in golang using the discordgo documentation. It answers comands like !js or !react with its documentation and recommended resources for learning
The bot was created with the aim to help members of a community I'm helping start in tech and for me to practice building apps in Go.

## How to use it

First you need to access the discord bot token and add it to a .env file, name it BOT_TOKEN.

Type and run the command:
```source .env```

So that the app knows where to look for the token

Then just build it with the ```go build main.go``` command and run it with the ```go run main.go``` command
