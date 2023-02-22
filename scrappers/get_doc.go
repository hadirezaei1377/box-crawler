package scrappers

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func get(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}
