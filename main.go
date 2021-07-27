package main

import (
	"topthema/bot"
	"topthema/crawler"
)

func main() {

	bot := bot.New()
	crawler := crawler.New()

	if result := crawler.Parse(); result != nil {
		bot.Post(*result)
	}

}
