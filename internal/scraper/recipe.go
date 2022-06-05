package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

type PageInfo struct {
	StatusCode int
	Text       string
}

func (s *Scraper) Scrape(url string) *PageInfo {

	c := colly.NewCollector()

	p := &PageInfo{}

	c.OnHTML(`script[type="application/ld+json"]`, func(h *colly.HTMLElement) {
		fmt.Printf("%s", h.Text)
		p.Text = h.Text
	})

	c.Visit(url)

	return p
}
