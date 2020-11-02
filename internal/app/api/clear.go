package api

import (
	logger "github.com/sanches1984/gopkg-logger"
	"net/http"
)

func Clear(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Debug(logger.App, "Clear")
	if client != nil {
		client.Clear()
	}
	w.WriteHeader(http.StatusOK)
}
