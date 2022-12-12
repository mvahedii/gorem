package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/word/create", app.createWord)
	mux.HandleFunc("/word/view", app.showWord)

	return mux
}
