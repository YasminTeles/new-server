//nolint:forbidigo
package settings

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YasminTeles/new-server/server"
)

func ListenAndServe() {
	server := server.NewServer()
	listeningPort := getListeningPort()

	printStartupServer()

	err := http.ListenAndServe(listeningPort, server)
	if err != nil {
		log.Fatalf("Error happened in listen and server. Err: %s", err)
	}
}

func getListeningPort() string {
	return fmt.Sprintf(":%s", Config.Port)
}

func printStartupServer() {
	fmt.Printf("Starting server on port %s...", Config.Port)
}
