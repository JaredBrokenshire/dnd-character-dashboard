package main

import (
	"dnd/api"
	"dnd/api/routes"
	"dnd/pkg/logging"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// @title D&D Character Dashboard
// @version 0.0.1
// @description A management solution for Dungeons & Dragons character sheet information

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @BasePath /
func main() {
	_ = godotenv.Load()

	server := api.NewServer()
	routes.RegisterRoutes(server)

	logging.Logger.LogInfo().Msg(server.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))).Error())
}
