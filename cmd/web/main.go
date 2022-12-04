package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/words", showAllWords)
	mux.HandleFunc("/word/create", createWord)
	mux.HandleFunc("/word/view", showWord)

	log.Print("Server Starting...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
