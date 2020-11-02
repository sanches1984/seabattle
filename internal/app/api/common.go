package api

import (
	"github.com/sanches1984/seabattle/internal/app/game"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

var client game.IClient

func prepareBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(model.NewError(err))
}
