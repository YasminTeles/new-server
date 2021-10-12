package main

import (
	"net/http"

	"github.com/YasminTeles/new-server/server"
	"github.com/YasminTeles/new-server/middleware"
	"github.com/urfave/negroni"
)

func main() {
	routers := server.Routers()

	n := negroni.Classic()

  n.Use(middleware.NewXRequestID())

	n.UseHandler(routers)

	http.ListenAndServe(":3000", n)
}
