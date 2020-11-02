package api

import (
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/game"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func CreateMatrix(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "CreateMatrix")
	request, err := model.NewCreateMatrixRequest(r)
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	client = game.NewGame(request.Range)
	w.WriteHeader(http.StatusOK)
}
