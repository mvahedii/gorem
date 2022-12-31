package handlers

import (
	"net/http"
)

func (server *HTTPServer) routes(wordHandler wordHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/word/create", wordHandler.createWord)
	mux.HandleFunc("/word/view", wordHandler.showWord)

	return mux
}
