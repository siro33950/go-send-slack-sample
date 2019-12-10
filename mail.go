package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[ERROR] Failed to loading .env file")
	}

	bot := NewSlackBot()
	if err := bot.SendMessage("z_bot-test", "fuga"); err != nil {
		log.Fatal("[ERROR] Failed to process send message")
	}
}

type Bot struct {
	Token string
}

func NewSlackBot() *Bot {
	var bot Bot
	if err := envconfig.Process("", &bot); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
	}
	return &bot
}

func (bot Bot) SendMessage(channel string, text string) (err error) {
	client := slack.New(bot.Token)

	params := slack.PostMessageParameters{
		Username: "hoge",
	}
	msgOptText := slack.MsgOptionText(text, true)
	msgOptParams := slack.MsgOptionPostMessageParameters(params)

	if _, _, err := client.PostMessage(channel, msgOptText, msgOptParams); err != nil {
		log.Printf("[ERROR] Failed to process post message: %s", err)
		return err
	}
	return nil
}
