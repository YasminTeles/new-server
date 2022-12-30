package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YasminTeles/new-server/middleware"
	"github.com/YasminTeles/new-server/settings"
	"github.com/urfave/negroni"
)

func NewServer() *http.Server {
	router := Router()

	handler := negroni.New()

	handler.Use(negroni.NewRecovery())
	handler.Use(middleware.NewXRequestID())
	handler.Use(middleware.NewLogger())

	handler.UseHandler(router)

	listeningPort := getListeningPort()

	server := &http.Server{
		Addr:         listeningPort,
		Handler:      handler,
		ReadTimeout:  settings.Config.Timeout,
		WriteTimeout: settings.Config.Timeout,
		IdleTimeout:  settings.Config.IdleTimeout,
	}

	return server
}

func getListeningPort() string {
	return fmt.Sprintf(":%s", settings.Config.Port)
}

func ListenAndServe() {
	server := NewServer()

	printStartupServer()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error happened in listen and server.")
	}
}

func printStartupServer() {
	fmt.Printf("Starting server on port %s...", settings.Config.Port)
}
