package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/mvahedii/gorem/internal/repositories"
	"github.com/mvahedii/gorem/internal/utils"
)

type HTTPServer struct {
	ErrLog  *log.Logger
	InfoLog *log.Logger
	Words   *repositories.WordModel
}

func NewHTTPServer(db *sql.DB, addr *string) *http.Server {

	httpServer := &HTTPServer{
		ErrLog:  utils.ErrLog,
		InfoLog: utils.InfoLog,
		Words:   &repositories.WordModel{DB: db},
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: utils.ErrLog,
		Handler:  httpServer.routes(),
	}

	return srv
}
