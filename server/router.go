package server

import (
	"fmt"
	"net/http"
)

var GitCommit string

func Routers() *http.ServeMux {
	routers := http.NewServeMux()

	routers.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Working!")
	})

	routers.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, GitCommit)
	})

	return routers
}
