package main

import (
	"github.com/YasminTeles/new-server/settings"
)

func main() {
	settings.LoadSettings()

	settings.ListenAndServe()
}
