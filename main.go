package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/fiber_1/database"
	"github.com/prranavv/fiber_1/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setuproutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.Getuser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	//Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.Updateproduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	//order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setuproutes(app)
	log.Fatal(app.Listen(":3000"))
}
