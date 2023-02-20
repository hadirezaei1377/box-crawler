package main

import (
	"box-crawler/database"
	"box-crawler/scrappers"
	"box-crawler/server"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly/v2"
)

func main() {
	db, err := database.New("mongodb://root:example@localhost:27017/default_db?authSource=admin")
	if err != nil {
		fmt.Println(err)
	}

	// initializing cron jobs
	//      get data
	//      process data and clean up data
	//      store data

	// for gocolly
c := colly.NewCollector(
	colly.AllowedDomains("box.live")
)

c.OnHTML("upcoming-fights-schedule","fight-results", func(e *colly.HTMLElement){
	metaTags := e.DOM.ParentsUntil("~").Find("meta")
	metaTags.Each(func(_ int , s *goquery.Selection){

	})
})


c.OnHTML("a[href]", func(e *colly.HTMLElement){
	link := e.Attw("href")
	c.Visit(r.Request.AbsouterURL(link))
})

c.OnRequest(func(r *colly.Request){
	fmt.Println("Crawl on page", r.URL.String())
})

c.Visit("https://box.live")

        scheduler := gocron.NewScheduler(time.UTC)
		scheduler.Every(3600).Second().Do(scrappers.ScrapeFightResults, db)







	/*
		c := colly.NewCollector(
			colly.AllowedDomains("https://box.live/fight-results"),
		)

		scheduler := gocron.NewScheduler(time.UTC)
		scheduler.Every(3600).Second().Do(scrappers.ScrapeFightResults, db)

		router := server.New(db)
		router.RegisterRoutes()
		if err := router.Start(); err != nil {
			log.Fatal(err)
		}

		c.OnCSS(".mw-parser-output", func(e *colly.CSSElement) {
			links := e.ChildAttrs("a", "href")
			fmt.Println(links)
		})
		c.Visit("https://box.live/fight-results/Web_scraping")
	*/
}
