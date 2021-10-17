package server

import (
	"github.com/YasminTeles/new-server/handlers"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/healthcheck", handlers.Healthcheck)
	router.GET("/version", handlers.Version)

	return router
}
