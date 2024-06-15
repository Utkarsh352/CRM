package routes

import (
	"github.com/gofiber/fiber/v2"

	"crm/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/customers", handlers.GetCustomers)
	app.Get("/customers/:id", handlers.GetCustomer)
	app.Post("/customers", handlers.CreateCustomer)
	app.Put("/customers/:id", handlers.UpdateCustomer)
	app.Delete("/customers/:id", handlers.DeleteCustomer)
}
