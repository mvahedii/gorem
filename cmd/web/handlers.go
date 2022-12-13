package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mvahedii/gorem/internal/models"
)

func (app *application) showWord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		app.errLog.Println(err)
		return
	}

	word, err := app.words.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", word)
}

func (app *application) createWord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	word := "dummy"
	description := "a model or replica of a human being."

	_, err := app.words.Insert(word, description)
	if err != nil {
		app.errLog.Fatal()
	}
}
