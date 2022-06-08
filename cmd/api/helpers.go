package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseWrapper map[string]interface{}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) json(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) error(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	rw := responseWrapper{"error": message}

	err := app.json(w, status, rw, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Not Found", r.Method)
	app.error(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s Method Not Allowed", r.Method)
	app.error(w, r, http.StatusMethodNotAllowed, message)
}
