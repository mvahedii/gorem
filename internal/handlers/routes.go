package handlers

import (
	"net/http"
)

func (server *HTTPServer) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/word/create", server.createWord)
	mux.HandleFunc("/word/view", server.showWord)

	return mux
}
