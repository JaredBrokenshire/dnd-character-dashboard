package main

import (
	"dnd/api"
	"dnd/api/routes"
	"dnd/pkg/logging"
	"dnd/pkg/validation"
	"fmt"
	"github.com/go-playground/validator/v10"
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

	logging.NewServerLogger()

	server := api.NewServer()
	server.Echo.Validator = validation.NewCustomValidator(validator.New())
	routes.RegisterRoutes(server)

	// Start Server
	logging.Logger.LogInfo().Msg(server.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))).Error())
}
