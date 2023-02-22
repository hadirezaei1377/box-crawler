package scrappers

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	_FR_FightDate      = "h3.site-content__section__subtitle:nth-child(1)"
	_FR_FightDateAlter = "h3.site-content__section__subtitle:nth-child(6)"
	_FR_FighterName1   = "header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)"
	_FR_FighterName2   = "header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)"
	_FighterScore1     = "header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2)"
	_FighterScore2     = "header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > div:nth-child(1)"
	_FR_FightInfo      = "div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)"
	_FR_MatchLocation  = "div:nth-child(2) > p:nth-child(1)"
	_ScoreCards        = "div:nth-child(2) > p:nth-child(2)"
)

func (sc scrapper) ScrapeFightResults() {
	doc, err := get("https://box.live/fight-results")
	if err != nil {
		fmt.Println("cannot get fight results: ", err)
	}

	selection := doc.Find("div.schedule-card > div").Each(func(i int, h *goquery.Selection) {
		fight_date := h.Find(_FightDate).Text()
		if fight_date == "" {
			fight_date = h.Find(_FightDateAlter).Text()
		}

		sc.db.InsertUpcomingFight(map[string]string{
			"fight_date":     fight_date,
			"fighter_name_1": h.Find(_FighterName1).Text(),
			"fighter_name_2": h.Find(_FighterName2).Text(),
			"fight_info":     h.Find(_FightInfo).Text(),
			"FighterScore1":  h.Find(_FighterScore1).Text(),
			"FighterScore2":  h.Find(_FighterScore2).Text(),
			"match_location": h.Find(_MatchLocation).Text(),
			"ScoreCards":     h.Find(_ScoreCards).Text(),
		})
	})

	fmt.Println(">>", selection)
}
