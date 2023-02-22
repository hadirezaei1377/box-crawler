package scrappers

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	_FightDate      = "h3.site-content__section__subtitle:nth-child(1)"
	_FightDateAlter = "h3.site-content__section__subtitle:nth-child(6)"
	_FighterName1   = "header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)"
	_FighterName2   = "header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)"
	_FightInfo      = "div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)"
	_MatchLocation  = "div:nth-child(2) > p:nth-child(1)"
)

func (sc scrapper) ScrapeUpcomingFights() {
	doc, err := get("https://box.live/upcoming-fights-schedule")
	if err != nil {
		fmt.Println("cannot get upcoming fights: ", err)
	}

	selection := doc.Find("#this-week > div").Each(func(i int, h *goquery.Selection) {
		fight_date := h.Find(_FightDate).Text()
		if fight_date == "" {
			fight_date = h.Find(_FightDateAlter).Text()
		}

		err := sc.db.InsertUpcomingFight(map[string]string{
			"fight_date":     fight_date,
			"fighter_name_1": h.Find(_FighterName1).Text(),
			"fighter_name_2": h.Find(_FighterName2).Text(),
			"fight_info":     h.Find(_FightInfo).Text(),
			"match_location": h.Find(_MatchLocation).Text(),
		})

		if err != nil {
			fmt.Println(err)
		}
	})

	fmt.Println(">>", selection)
}
