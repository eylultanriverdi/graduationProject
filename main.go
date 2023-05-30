package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)
	app.Listen("localhost:3000")

}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Post("/register", api.RegisterHandler)
	app.Post("/product", api.ProductHandler)
	app.Get("/products", api.GetProducts)
	return app
}
