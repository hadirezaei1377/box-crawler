package scrappers

import "box-crawler/database"

func ScrapeUpcomingFights(db database.Database) []string {
	return nil
}
func ScrapeFightResults(db database.Database) {

	http.Get("https://box.live/fight-results")

	fight := entities.Fight{}
	db.UpsertUpcomingFights(fight)
}
