package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func showAllWords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from show all words"))
}

func showWord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "the id is %d", id)
}

func createWord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Hello from create word"))
}
