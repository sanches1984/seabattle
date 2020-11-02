package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"github.com/sanches1984/seabattle/internal/app/model"
	"net/http"
)

func Shot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Shot")
	if client == nil {
		prepareBadRequest(w, fmt.Errorf("Game not created"))
		return
	}
	if !client.IsStarted() {
		prepareBadRequest(w, fmt.Errorf("Game not started or already ended"))
		return
	}
	request, err := model.NewShotRequest(r)
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	response, err := client.MakeShot(request.GetShot())
	if err != nil {
		prepareBadRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response.GetJSON())
}
