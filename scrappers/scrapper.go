package scrappers

import (
	"box-crawler/database"
)

type scrapper struct {
	db database.Database
}

func New(db database.Database) *scrapper {
	return &scrapper{
		db: db,
	}
}
