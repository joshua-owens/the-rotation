package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RecipeImport struct {
	Url string `json:"url"`
}

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Post("/recipes/import", func(w http.ResponseWriter, r *http.Request) {
		var ri RecipeImport
		err := json.NewDecoder(r.Body).Decode(&ri)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		app.logger.Println(ri.Url)
	})

	return r
}
