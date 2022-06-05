package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ImportRequest struct {
	Url string `json:"url"`
}

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"hello": "world",
		}
		app.json(w, http.StatusOK, data, nil)
	})

	r.Post("/recipes/import", func(w http.ResponseWriter, r *http.Request) {
		var ir ImportRequest
		err := json.NewDecoder(r.Body).Decode(&ir)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		app.scraper.Scrape(ir.Url)
	})

	return r
}
