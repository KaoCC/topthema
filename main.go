package main

import (
	"log"
	"topthema/bot"
	"topthema/crawler"
)

func main() {

	bot := bot.New()
	crawler := crawler.New()

	if result := crawler.Parse(); result != nil {
		bot.Post(*result)
	} else {
		log.Println("No update from source")
	}

	crawler.Save()

}
