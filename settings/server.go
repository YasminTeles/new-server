//nolint:forbidigo
package settings

import (
	"fmt"
	"net/http"

	"github.com/YasminTeles/new-server/server"
	log "github.com/sirupsen/logrus"
)

func ListenAndServe() {
	server := server.NewServer()
	listeningPort := getListeningPort()

	printStartupServer()

	err := http.ListenAndServe(listeningPort, server)
	if err != nil {
		log.Fatal("Error happened in listen and server.")
	}
}

func getListeningPort() string {
	return fmt.Sprintf(":%s", Config.Port)
}

func printStartupServer() {
	fmt.Printf("Starting server on port %s...", Config.Port)
}
