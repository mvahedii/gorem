package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var ErrLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ErrLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
