package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mvahedii/gorem/internal/repositories"
	"github.com/mvahedii/gorem/internal/services"
	"github.com/mvahedii/gorem/internal/utils"
)

func (httpServer *HTTPServer) showWord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		utils.NotFound(w)
		httpServer.ErrLog.Println(err)
		return
	}

	word, err := services.ShowWordService(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNoRecord) {
			utils.NotFound(w)
		} else {
			utils.ServerError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", word)
}

func (httpServer *HTTPServer) createWord(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		utils.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		utils.ClientError(w, http.StatusBadRequest)
		return
	}
	word := r.PostForm.Get("word")
	description := r.PostForm.Get("description")
	fmt.Println(word, description)
	_, err = services.CreateWordService(word, description)
	if err != nil {
		httpServer.ErrLog.Fatal()
	}
}
