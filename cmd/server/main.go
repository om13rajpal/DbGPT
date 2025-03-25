package main

import (
	"github.com/om13rajpal/dbgpt/config"
	"github.com/om13rajpal/dbgpt/internal/database"
	"github.com/om13rajpal/dbgpt/internal/routes"
)

func main() {
	config.InitConfig()
	database.ConnectMongo()

	router := routes.InitRoute()

	router.Run(":" + config.PORT)
}
