package data

import "time"

type Recipe struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Time         int32     `json:"time"`
	Servings     int32     `json:"servings"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	Image        string    `json:"image"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
