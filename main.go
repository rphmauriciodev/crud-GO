package main

import (
	"github.com/rphmauriciodev/crud-GO.git/config"
	"github.com/rphmauriciodev/crud-GO.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}
