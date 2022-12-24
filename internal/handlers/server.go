package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/mvahedii/gorem/internal/repositories"
	"github.com/mvahedii/gorem/internal/utils"
)

type HTTPServer struct {
	errLog  *log.Logger
	infoLog *log.Logger
	words   *repositories.WordModel
}

func NewHTTPServer(db *sql.DB, addr *string) *http.Server {

	httpServer := &HTTPServer{
		errLog:  utils.ErrLog,
		infoLog: utils.InfoLog,
		words:   &repositories.WordModel{DB: db},
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: utils.ErrLog,
		Handler:  httpServer.routes(),
	}

	return srv
}
