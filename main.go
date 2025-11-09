package main

// @title go.learnings.api.rest by h0wzy
// @version 1.0
// @description This is a sample server for a job openings management API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

import (
	"github.com/H0wZy/go.learnings.api.rest/config"
	_ "github.com/H0wZy/go.learnings.api.rest/docs"
	"github.com/H0wZy/go.learnings.api.rest/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitDb()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	router.Initialize()
}
