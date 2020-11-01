package api

import (
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/game"
	"net/http"
)

func Clear(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Clear")
	game.ClearGame()
	w.WriteHeader(http.StatusOK)
}
