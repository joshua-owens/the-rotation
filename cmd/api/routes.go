package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"hello": "world",
		}
		app.writeJSON(w, http.StatusOK, data, nil)
	})

	r.Post("/recipes/import", app.importRecipeHandler)

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	return r
}
