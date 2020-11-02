package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func Ship(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Ship")
	if client == nil {
		prepareBadRequest(w, fmt.Errorf("Game not created"))
		return
	}
	if client.IsStarted() {
		prepareBadRequest(w, fmt.Errorf("Game already started"))
		return
	}

	request, err := model.NewShipRequest(r)
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	ships, err := request.GetShips()
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	err = client.CreateShips(ships)
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
