package api

import (
	"github.com/sanches1984/seabattle/internal/app/game"
	"net/http"
)

func GameInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(game.GetInfo()))
}
