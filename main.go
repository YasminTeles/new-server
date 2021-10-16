package main

import (
	"fmt"
	"net/http"

	"github.com/YasminTeles/new-server/server"
)

func main() {
	server := server.NewServer()

	fmt.Println(`Starting server on port 3000...`)

	http.ListenAndServe(":3000", server)
}
