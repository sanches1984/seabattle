package app

import (
	"context"
	logger "github.com/sanches1984/gopkg-logger"
	"net/http"
)

type App struct {
	url string
}

func NewApp(url string, handlers map[string]func(w http.ResponseWriter, r *http.Request)) *App {
	for path, handler := range handlers {
		http.HandleFunc(path, handler)
	}
	return &App{url: url}
}

func (a *App) Run(ctx context.Context) error {
	logger.Info(ctx, "Start listening %s", a.url)
	if err := http.ListenAndServe(a.url, nil); err != nil {
		return err
	}
	return nil
}
