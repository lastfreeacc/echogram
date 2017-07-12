package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/lastfreeacc/echogram/teleapi"
)

const (
	confFilename = "echogram.conf.json"
)

var (
	conf = make(map[string]interface{})
	botToken string
	bot teleapi.Bot
)

func main() {
	myInit()
	upCh := bot.Listen()
	for update := range upCh {
		bot.SendMessage(update.Message.Chat.ID, update.Message.Text)
	}
}

func myInit() {
	readMapFromJSON(confFilename, &conf)
	botToken, ok := conf["botToken"]
	if !ok || botToken == "" {
		log.Fatalf("[Error] can not find botToken in config file: %s\n", confFilename)
	}
	bot = teleapi.NewBot(botToken.(string))
}

func readMapFromJSON(filename string, mapVar *map[string]interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("[Warning] can not read file '%s'\n", filename)
	}
	if err := json.Unmarshal(data, mapVar); err != nil {
		log.Fatalf("[Warning] can not unmarshal json from file '%s'\n", filename)
	}
	log.Printf("[Info] read data from file: %s:\n%v\n", filename, mapVar)
}