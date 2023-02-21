package scrappers

import (
	"box-crawler/database"
	"box-crawler/entities"
	"net/http"
)

func ScrapeFightResults(db database.Database) {

	http.Get("https://box.live/fight-results")

	fight := entities.Fight{}
	db.UpsertFightResults(fight)
}
