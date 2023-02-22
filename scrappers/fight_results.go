package scrappers

func (sc scrapper) ScrapeFightResults() {
	get("https://box.live/fight-results")

	// d.OnHTML("#latest-results", func(container *colly.HTMLElement) {
	// 	container.ForEach("div", func(i int, h *colly.HTMLElement) {
	// 		fightDate := h.ChildText("?")
	// 		fighterName1 := h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(1)")
	// 		fighterName2 := h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > span:nth-child(2)")
	// 		fighterScore1 := h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(1) > div:nth-child(2)")
	// 		fighterScore2 := h.ChildText("header:nth-child(1) > div:nth-child(3) > div:nth-child(3) > div:nth-child(1) > div:nth-child(1)")
	// 		fightInfo := h.ChildText("div.schedule-card:nth-child(7) > header:nth-child(1) > div:nth-child(5) > a:nth-child(1)")
	// 		matchLocation := h.ChildText("div:nth-child(2) > p:nth-child(1)")
	// 		scoreCards := h.ChildText("div:nth-child(2) > p:nth-child(2)")
	// 		fmt.Println(fightDate, fighterName1, fighterName2, fighterScore1, fighterScore2, fightInfo, matchLocation, scoreCards)
	// 	})
	// })

	// sc.db.UpsertFightResults(map[string]string{})
}
