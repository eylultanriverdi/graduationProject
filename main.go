package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)
	app.Listen("localhost:3001")

}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Post("/userRegister", api.RegisterHandler)
	app.Post("/nutritionistRegister", api.NutritionistRegisterHandler)
	app.Post("/product", api.ProductHandler)
	app.Post("/signin", api.SigninHandler)
	app.Post("/signinNutritionist", api.SigninNutritionistHandler)
	app.Get("/profile", api.ProfileHandler)
	app.Get("/profileNutritionist", api.ProfileNutritionistHandler)
	app.Post("/dietCategory", api.DietCategoryHandler)
	app.Post("/addList", api.HandleAddListProduct)
	app.Post("/addNutritionistList", api.HandleAddNutritionistList)
	app.Get("/nutritionistList", api.GetNutritionistList)
	app.Get("/products", api.GetProducts)
	app.Get("/nutritionists", api.GetNutritionists)
	app.Get("/dietCategories", api.GetDietCategories)
	app.Get("/calorieInfo", api.GetCalorieList)
	return app
}
