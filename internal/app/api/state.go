package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/game"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func State(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "State")
	if !game.IsGameStarted() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(fmt.Errorf("Game not started")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(game.GetStat().GetJSON())
}
