package scraper

import (
	"encoding/json"

	"github.com/gocolly/colly/v2"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

type ListElement struct {
	SchemaType string `json:"@type"`
	SchemaItem string `json:""`
}

type Schema struct {
	ListElements []ListElement
}

func (s *Scraper) Scrape(url string) Schema {

	c := colly.NewCollector()

	schema := *&Schema{}

	c.OnHTML(`script[type="application/ld+json"]`, func(h *colly.HTMLElement) {
		json.Unmarshal([]byte(h.Text), &schema.ListElements)
	})

	c.Visit(url)

	return schema
}
