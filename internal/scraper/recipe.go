package scraper

import (
	"fmt"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape(url string) {
	fmt.Printf("%s", url)
}
