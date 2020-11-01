package main

import (
	"context"
	"github.com/sanches1984/seabattle/internal/app/api"
	"github.com/sanches1984/seabattle/internal/config"
	"github.com/sanches1984/seabattle/internal/pkg/app"
	"net/http"
)

var handlers = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":              api.GameInfo,
	"/create-matrix": api.CreateMatrix,
	"/ship":          api.Ship,
	"/shot":          api.Shot,
	"/clear":         api.Clear,
	"/state":         api.State,
}

func main() {
	config.Load()
	ctx := context.Background()
	a := app.NewApp(config.Host(), handlers)
	err := a.Run(ctx)
	if err != nil {
		panic(err)
	}
}
