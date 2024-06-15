package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"crm/database"
	"crm/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	database.ConnectDB()
	defer database.DisconnectDB()

	routes.SetupRoutes(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on port 8080")
}
