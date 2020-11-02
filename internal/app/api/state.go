package api

import (
	"fmt"
	logger "github.com/sanches1984/gopkg-logger"
	"net/http"
)

func State(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "State")
	if client == nil {
		prepareBadRequest(w, fmt.Errorf("Game not created"))
		return
	}
	if !client.IsStarted() {
		prepareBadRequest(w, fmt.Errorf("Game not started"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(client.GetStat().GetJSON())
}
