package main

import (
	"github.com/om13rajpal/dbgpt/config"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/routes"
)

func main() {
	config.InitConfig()
	database.ConnectPostgres()

	router := routes.InitRoute()

	router.Run(":" + config.PORT)

	defer database.Pool.Close()
}
