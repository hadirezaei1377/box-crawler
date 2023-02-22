package scrappers

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func (sc scrapper) ScrapeUpcomingFights() {
	doc, err := get("https://box.live/upcoming-fights-schedule")
	if err != nil {
		log.Fatal(">> ", err)
	}

	fmt.Println(doc.Find(_Date).Text())

	selection := doc.Find("#this-week > div").Each(func(i int, h *goquery.Selection) {
		fight_date := h.Find("h3.site-content__section__subtitle:nth-child(1)").Text()
		if fight_date == "" {
			fight_date = h.Find("h3.site-content__section__subtitle:nth-child(6)").Text()
		}
		fighter_name1 := h.Find("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)").Text()
		fighterName2 := h.Find("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)").Text()
		fightInfo := h.Find("div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)").Text()
		matchLocation := h.Find("div:nth-child(2) > p:nth-child(1)").Text()
		fmt.Println(fighter_name1, fighterName2, fightInfo, matchLocation)

	})
	fmt.Println(">>", selection)
}
