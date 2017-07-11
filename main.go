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

func main() {
	conf := make(map[string]interface{})
	readMapFromJSON(confFilename, &conf)
	botToken, ok := conf["botToken"]
	if !ok || botToken == "" {
		log.Fatalf("[Error] can not find botToken in config file: %s\n", confFilename)
	}
	bot := teleapi.NewBot(botToken.(string))
	bot.SendMessage(123456, "chat text")
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