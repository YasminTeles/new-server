package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	GitCommit string
	GitTag    string
	BuildData string
)

func Routers() *http.ServeMux {
	routers := http.NewServeMux()

	routers.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Working!")
	})

	routers.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)
		resp["version"] = GitTag
		resp["build"] = strings.Replace(BuildData, "\t", " ", 1)
		resp["commit"] = GitCommit

		json.NewEncoder(w).Encode(resp)
	})

	return routers
}
