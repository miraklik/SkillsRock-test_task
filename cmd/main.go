package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miraklik/TODO-list/config"
	"github.com/miraklik/TODO-list/db"
	"github.com/miraklik/TODO-list/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to get .env: %v", err)
	}

	handlers, err := handler.NewTaskHandler()
	if err != nil {
		log.Fatalf("Failed to create task handler: %v", err)
	}

	if err := db.InitSchema(); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	router := fiber.New()

	router.Post("/tasks", handlers.CreateTask)
	router.Get("/tasks", handlers.GetAllTasks)
	router.Put("/tasks/:id", handlers.UpdateTaks)
	router.Delete("tasks/:id", handlers.DeleteTask)

	if err := router.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
