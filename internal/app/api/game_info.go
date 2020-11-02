package api

import (
	"net/http"
)

func GameInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	if client != nil {
		w.Write([]byte(client.GetInfo()))
	}
}
