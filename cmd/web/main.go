package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Port")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/words", app.showAllWords)
	mux.HandleFunc("/word/create", app.createWord)
	mux.HandleFunc("/word/view", app.showWord)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Print("Server Starting...", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errLog.Fatal(err)
	}
}
