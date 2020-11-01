package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/game"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func Shot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Shot")
	if !game.IsGameStarted() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(fmt.Errorf("Game not started or already ended")))
		return
	}
	request, err := model.NewShotRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(err))
		return
	}

	response, err := game.MakeShot(request.GetShot())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response.GetJSON())
}
