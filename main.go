package main

import (
	"github.com/H0wZy/go.learnings.restapi/config"
	"github.com/H0wZy/go.learnings.restapi/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitDb()
	if err != nil {
		logger.Infof("config init error: %v", err)
		return
	}
	router.Initialize()
}
