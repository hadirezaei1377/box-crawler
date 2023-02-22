package main

import (
	"box-crawler/database"
	"box-crawler/scrappers"
	"box-crawler/server"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	db, err := database.New("mongodb://root:example@localhost:27017/default_db?authSource=admin")
	if err != nil {
		panic(err)
	}

	scrapper := scrappers.New(db)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(2).Second().Do(scrapper.ScrapeUpcomingFights)
	scheduler.Every(2).Second().Do(scrapper.ScrapeFightResults)
	scheduler.StartAsync()

	router := server.New(db)
	router.RegisterRoutes()
	if err := router.Start(); err != nil {
		log.Fatal(err)
	}
}
