package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) importRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Url string `writeJSON:"url"`
	}
	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
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
