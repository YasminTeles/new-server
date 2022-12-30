//nolint:gochecknoinits,exhaustivestruct
package main

import (
	"github.com/YasminTeles/new-server/server"
	"github.com/YasminTeles/new-server/settings"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
}

func main() {
	settings.LoadSettings()

	server.ListenAndServe()
}
