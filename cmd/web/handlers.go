package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) showAllWords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from show all words"))
}

func (app *application) showWord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		app.errLog.Println(err)
		return
	}
	fmt.Fprintf(w, "the id is %d", id)
}

func (app *application) createWord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Hello from create word"))
}
