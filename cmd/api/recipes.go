package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) importRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	p := app.scraper.Scrape(input.Url)

	b, err := json.Marshal(p)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

}
