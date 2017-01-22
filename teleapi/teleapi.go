package teleapi

import (
	"bytes"
	"fmt"
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

type teleBot struct {
	token string
}

func (bot *teleBot) makeURL(m method) string {
	return fmt.Sprintf("%s%s/%s", botURL, bot.token, m)
}

// NewBot ...
func NewBot(t string) *teleBot {
	bot := teleBot{token: t}
	return &bot
}

// SendMessage ...
func (bot *teleBot) SendMessage(chatID string, text string) error {
	jsonStr := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, chatID, text)
	json := []byte(jsonStr)
	endPnt := bot.makeURL(sendMessageMthd)
	req, err := http.NewRequest("POST", endPnt, bytes.NewBuffer(json))
	if err != nil {
		log.Printf("{Error} in build req: %s", err.Error())
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
	// bosy, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("[Warning] can not read api answer: {method: %s, data:%s}, err: %s", sendMessageMthd, json, err)
	// }
	return nil
}
