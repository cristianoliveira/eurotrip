package common

import "os"

type Setting struct {
	Port     string
	FilePath string
}

func Settings() *Setting {
	s := &Setting{
		Port:     "8088",
		FilePath: "../../../../../data/example",
	}

	port := os.Getenv("PORT")
	if len(port) != 0 {
		s.Port = port
	}

	filePath := os.Getenv("FILEPATH")
	if len(filePath) != 0 {
		s.FilePath = filePath
	}

	return s
}
