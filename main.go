package main

import (
	"github.com/lastfreeacc/teleautomonbot/teleapi"
)

func main() {
	bot := teleapi.NewBot("token")
	bot.SendMessage("@chat", "some test")
}