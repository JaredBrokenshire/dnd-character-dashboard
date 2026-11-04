package main

import (
	"dnd-character-dashboard/api/routes"
	_ "dnd-character-dashboard/docs"
	"dnd-character-dashboard/pkg/logging"
	"dnd-character-dashboard/pkg/validation"
	"github.com/go-playground/validator/v10"

	"dnd-character-dashboard/api"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// @title D&D Character Dashboard
// @version 0.1.1
// @description A management solution for Dungeons & Dragons character sheet information

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @BasePath /
func main() {
	// Load .env when present (ignored in Docker where env vars are injected directly).
	_ = godotenv.Load()

	logging.NewServerLogger()

	server := api.NewServer()
	server.Echo.Validator = validation.NewCustomValidator(validator.New())

	routes.RegisterRoutes(server)

	// ── Start ──────────────────────────────────────────────────────────────────
	logging.Logger.LogInfo().Msg(server.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))).Error())
}
