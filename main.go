package main

import (
	"github.com/lastfreeacc/echogram/teleapi"
)

func main() {
	bot := teleapi.NewBot("token")
	bot.SendMessage(123456, "chat text")
}