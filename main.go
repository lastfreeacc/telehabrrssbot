package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lastfreeacc/telehabrrssbot/rss"
	bot "github.com/lastfreeacc/telehabrrssbot/teleapi"
)

const botToken = "botToken"

var guids = make(map[string]interface{}) // map[string]bool
var guidsFilename = "guids.json"
var conf = make(map[string]interface{}) // map[string]string
var confFilename = "conf.json"

func main() {
	done := make(chan bool)
	feeds := make(chan rss.Item, 20)
	go readHabrRss(feeds)
	go sendTeleMessage(feeds)
	<-done
}

func readHabrRss(c chan<- rss.Item) {
	habrURL := "https://habrahabr.ru/rss/hub/go/"
	for {
		feed, err := rss.ReadRssURL(habrURL)
		if err != nil {
			log.Printf("[Warning] some problems: %s\n", err.Error())
			time.Sleep(time.Minute)
			continue
		}
		items := feed.Channel.Items
		log.Printf("[Info] read %d items form %s\n", len(items), habrURL)
		for _, item := range items {
			c <- item
		}
		time.Sleep(10 * time.Minute)
	}
}

func sendTeleMessage(c <-chan rss.Item) {
	token := conf[botToken].(string)
	log.Printf("[Info] bot token: %s\n", token)
	bot := bot.NewBot(token)
	for {
		item := <-c
		guid := item.GUID
		if _, ok := guids[guid]; ok {
			continue
		}
		guids[guid] = true
		text := fmt.Sprintf("%s\n%s", item.Title, guid)
		bot.SendMessage("@saska_me", text)
		log.Printf("[Info] sent message to tele: %s", guid)
		saveGuids(guids)
	}
}

func init() {
	readMapFromJSON(guidsFilename, &guids)
	readMapFromJSON(confFilename, &conf)
	token, ok := conf[botToken]
	if !ok || token == "" {
		log.Fatalf("[Error] can not find botToken in config file\n")
	}
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

func saveGuids(guids map[string]interface{}) {
	data, err := json.Marshal(guids)
	if err != nil {
		log.Printf("[Warning] can not marshal guids: %s", err.Error())
		return
	}
	file, err := os.OpenFile(guidsFilename, os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[Warning] can not open file: %s, error: %s", guidsFilename, err.Error())
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Printf("[Warning] can not write file: %s", err.Error())
	}
	err = file.Sync()
	if err != nil {
		log.Printf("[Warning] can not sync data to file: %s", err.Error())
	}
	log.Printf("[Info] guids successfully saved")
}
