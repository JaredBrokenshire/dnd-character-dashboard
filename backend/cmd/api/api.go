package main

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/pkg/validation"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"os"
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

	server := api.NewServer()
	server.Echo.Validator = validation.NewCustomValidator(validator.New())

	if err := server.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))).Error(); err != "" {
		panic(err)
	}
}
