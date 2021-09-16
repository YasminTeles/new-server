package server

import (
	"fmt"
	"net/http"
)

func Routers() *http.ServeMux {
	routers := http.NewServeMux()

	routers.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	return routers
}
