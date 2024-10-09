package main

import (
	"github.com/vitorsoratto/gojobs/config"
	"github.com/vitorsoratto/gojobs/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger()

	err := config.InitDatabase()
	if err != nil {
		logger.Errorf("Error initializing the database: %s", err)
		return
	}

	router.InitRouter()
}
