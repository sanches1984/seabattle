package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/game"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func Ship(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Ship")
	if game.IsGameStarted() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(fmt.Errorf("Game already started")))
		return
	}

	request, err := model.NewShipRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(err))
		return
	}

	ships, err := request.GetShips()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(err))
		return
	}

	err = game.CreateShips(ships)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(model.NewError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
