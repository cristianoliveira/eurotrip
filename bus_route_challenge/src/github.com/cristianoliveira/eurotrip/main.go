package main

import (
	"github.com/cristianoliveira/eurotrip/api"
	"github.com/cristianoliveira/eurotrip/common"
	"os"
)

func main() {
	settings := common.Settings()

	filePath := os.Args[1]
	if len(filePath) != 0 {
		settings.FilePath = filePath
	}

	api.Serve(settings)
}
