package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"tamthanh/golang-myapp/database"
	"tamthanh/golang-myapp/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
}
func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":7000"))
}
