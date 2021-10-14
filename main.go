package main

import (
	"net/http"

	"github.com/YasminTeles/new-server/middleware"
	"github.com/YasminTeles/new-server/server"
	"github.com/urfave/negroni"
)

func main() {
	routers := server.Routers()

	n := negroni.New()

	n.Use(negroni.NewRecovery())
	n.Use(middleware.NewXRequestID())
	n.Use(middleware.NewLogger())

	n.UseHandler(routers)

	http.ListenAndServe(":3000", n)
}
