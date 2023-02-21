package scrappers

import (
	"box-crawler/database"
	"box-crawler/entities"
)

func ScrapeUpcomingFights(db database.Database) []string {
	return nil
}
func ScrapeFightResults(db database.Database) {

	fight := entities.Fight{}
	db.UpsertUpcomingFights(fight)
}
