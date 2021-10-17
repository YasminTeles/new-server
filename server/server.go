package server

import (
	"github.com/YasminTeles/new-server/middleware"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	router := Router()

	server := negroni.New()

	server.Use(negroni.NewRecovery())
	server.Use(middleware.NewXRequestID())
	server.Use(middleware.NewLogger())

	server.UseHandler(router)

	return server
}
