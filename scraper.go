package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

const (
	URLBase = "https://www.bbc.co.uk"
	URL     = URLBase + "/sounds/search?q=news"
)

var (
	debug          bool
	foundSomething bool
)

func main() {
	parseArgs()

	if debug {
		log.Println("Starting scraper")
	}

	episodes := []*colly.HTMLElement{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0"),
	)
	c.Limit(&colly.LimitRule{
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		if debug {
			log.Println("Visiting", r.URL)
		}
	})

	c.OnHTML("ul.sc-c-list__items li article a", func(e *colly.HTMLElement) {
		episodes = append(episodes, e)
	})

	c.OnXML("//ul", func(e *colly.XMLElement) {
		if len(episodes) > 0 {
			foundSomething = true
		}
	})

	c.OnScraped(func(r *colly.Response) {
		if debug {
			log.Println("Finished", r.Request.URL)
		}
		if foundSomething {
			fmt.Println(URLBase + episodes[0].Attr("href"))
		}
	})

	err := c.Visit(URL)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func parseArgs() {
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		debug = true
	}
}
