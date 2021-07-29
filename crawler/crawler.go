package crawler

import (
	"fmt"
	"log"
	"time"
	"topthema/util"

	"github.com/mmcdole/gofeed"
)

type crawler struct {
	parser     *gofeed.Parser
	lastUpdate time.Time
}

const dwURL string = "https://rss.dw.com/xml/DKpodcast_topthemamitvokabeln_de"
const filename string = "last.txt"

func New() *crawler {
	fp := gofeed.NewParser()

	return &crawler{
		parser:     fp,
		lastUpdate: util.GetLastTime(filename),
	}
}

/// This function returns nil if there is no update.
func (crawler *crawler) Parse() *util.Record {

	feed, err := crawler.parser.ParseURL(dwURL)
	if err != nil {
		log.Printf("%s\n", err)
		return nil
	}

	fmt.Println(feed.PublishedParsed)

	if feed.PublishedParsed.After(crawler.lastUpdate) {

		crawler.lastUpdate = *feed.PublishedParsed

		item := feed.Items[0]

		return &util.Record{
			Title: item.Title,
			Link:  item.Link,
		}

	}

	return nil
}

/// Save the state of the crawler.
/// Currently, only the last update time is saved.
func (crawler *crawler) Save() {
	util.SetLastTime(filename, crawler.lastUpdate)
}
