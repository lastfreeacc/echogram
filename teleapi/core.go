package teleapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type method string

// const sendMessage method = method("sendMessage")

// type telegramBot struct {
// 	token string
// }
//
// func (bot *telegramBot) getURL() string {
// 	url := "https://api.telegram.org/bot%s"
// 	return fmt.Sprintf(url, bot.token)
// }

type method string

const botURL string = "https://api.telegram.org/bot"
const sendMessageMthd method = "sendMessage"
const getUpdates method = "getUpdates"

type bot struct {
	token        string
	updateCh     chan *Update
	currenOffset int64
}

func (bot *bot) makeURL(m method) string {
	return fmt.Sprintf("%s%s/%s", botURL, bot.token, m)
}

// Bot ...
type Bot interface {
	SendMessage(int64, string) error
}

// NewBot ...
func NewBot(t string) Bot {
	bot := bot{
		token:        t,
		updateCh:     make(chan *Update, 100),
		currenOffset: 0,
	}
	return &bot
}

// SendMessage ...
func (bot *bot) SendMessage(chatID int64, text string) error {
	jsonStr := fmt.Sprintf(`{"chat_id":"%d","text":"%s"}`, chatID, text)
	json := []byte(jsonStr)
	endPnt := bot.makeURL(sendMessageMthd)
	req, err := http.NewRequest("POST", endPnt, bytes.NewBuffer(json))
	if err != nil {
		log.Printf("[Error] in build req: %s", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Error] in send req: %s", err.Error())
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Warning] can not read api answer: {method: %s, data:%s}, err: %s", sendMessageMthd, json, err)
	}
	return nil
}