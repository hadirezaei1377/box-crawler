package main

import (
	"box-crawler/database"
	"box-crawler/scrappers"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
)

func main() {
	db, err := database.New("mongodb://root:example@localhost:27017/default_db?authSource=admin")
	if err != nil {
		fmt.Println(err)
	}

	fightsCollection := client.Database("testing").Collection("fights")
	fight := bson.D{{"fight-results", "results"}, {"upcoming-fights", "upcoming-fights"}}
	result, err := fightsCollection.InsertOne(context.TODO(), fight)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)

	fights := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}

	results, err := fightsCollection.InsertMany(context.TODO(), fights)

	if err != nil {
		panic(err)
	}

	fmt.Println(results.InsertedIDs)

	c := colly.NewCollector(
		colly.AllowedDomains("box.live"),
		colly.CacheDir("./box-live-cache"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawl on page", r.URL.String())
	})


	// upcoming-fights-schedule
	c.Visit("https://box.live/upcoming-fights-schedule")

	
	c.OnHTML("#this-week", func(container *colly.HTMLElement) {
		container.ForEach("div", func(i int, h *colly.HTMLElement) {
			 
            fmt.Println("FightDate:     ", h.ChildText("h3.site-content__section__subtitle:nth-child(اعداد فرد )"))
			fmt.Println("FighterName1:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)"))
            fmt.Println("FighterName2:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)"))
            fmt.Println("FightInfo:     ", h.ChildText("div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)"))
            fmt.Println("MatchLocation: ", h.ChildText("div:nth-child(2) > p:nth-child(1)"))
			
			})
		})

		

		// fight result 
		d.Visit("https://box.live/fight-results/")

		d.OnHTMLc.OnHTML("#latest-results", func(container *colly.HTMLElement) {
			container.ForEach("div", func(i int, h *colly.HTMLElement) {
				 
				fmt.Println("FightDate:     ", h.ChildText("?"))
				fmt.Println("FighterName1:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)"))
				fmt.Println("FighterName2:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)"))
				fmt.Println("FighterScore1:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2)"))
				fmt.Println("FighterScore2:  ", h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > div:nth-child(1)"))
				fmt.Println("FightInfo:     ", h.ChildText("div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)"))
				fmt.Println("MatchLocation: ", h.ChildText("div:nth-child(2) > p:nth-child(1)"))
				fmt.Println("ScoreCards:    ", h.ChildText("div:nth-child(2) > p:nth-child(2)"))
			})
		})
		
		


	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(3600).Second().Do(scrappers.ScrapeFightResults, db)

	func hello(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	  }

	router := server.New(db)
	router.RegisterRoutes()
	if err := router.Start(); err != nil {
		log.Fatal(err)
	}

	fightsCollection := client.Database("testing").Collection("fights")
	fight := bson.D{{"fight-results", "results"}, {"upcoming-fights", "upcoming-fights"}}
	result, err := fightsCollection.InsertOne(context.TODO(), fight)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)

	results, err := fightsCollection.InsertMany(context.TODO(), fights)

	if err != nil {
		panic(err)
	}

	fmt.Println(results.InsertedIDs)

}