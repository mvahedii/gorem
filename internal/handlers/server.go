package handlers

import (
	"log"
	"net/http"

	"github.com/mvahedii/gorem/internal/utils"
)

type HTTPServer struct {
	ErrLog  *log.Logger
	InfoLog *log.Logger
}

func NewHTTPServer(wordHandler wordHandler, addr *string) *http.Server {

	httpServer := &HTTPServer{
		ErrLog:  utils.ErrLog,
		InfoLog: utils.InfoLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: utils.ErrLog,
		Handler:  httpServer.routes(wordHandler),
	}

	return srv
}
