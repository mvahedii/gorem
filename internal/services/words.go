package server

import (
	"github.com/mvahedii/gorem/internal/handlers"
	"github.com/mvahedii/gorem/internal/repositories"
)

func ShowWordService(httpServer *handlers.HTTPServer, id int) (*repositories.Word, error) {
	word, err := httpServer.Words.Get(id)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func CreateWordService(httpServer *handlers.HTTPServer, word, description string) (int, error) {
	status, err := httpServer.Words.Insert(word, description)
	if err != nil {
		return status, err
	}
	return status, nil
}
