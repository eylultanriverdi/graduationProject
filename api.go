package main

import (
	"encoding/base64"

	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) RegisterHandler(c *fiber.Ctx) error {
	register := models.RegisterDTO{}
	err := c.BodyParser(&register)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	createUser, err := api.Service.Register(register)

	switch err {
	case nil:
		return c.JSON(createUser)
	case UserAlreadyExistError, PasswordHashingError:
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}

func (api *Api) ProductHandler(c *fiber.Ctx) error {
	product := models.ProductCategoryDTO{}
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Decode base64 string to byte array
	productImage, err := base64.StdEncoding.DecodeString(product.ProductImage)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Create Product instance with byte array
	createProduct, err := api.Service.GetProduct(models.ProductCategoryDTO{
		ProductName:       product.ProductName,
		Description:       product.Description,
		ProductImage:      string(productImage),
		ProteinValue:      product.ProteinValue,
		CarbohydrateValue: product.CarbohydrateValue,
		OilValue:          product.OilValue,
		GlutenValue:       product.GlutenValue,
		KetogenicDiet:     product.KetogenicDiet,
		GlutenFree:        product.GlutenFree,
		SaltFree:          product.SaltFree,
	})

	switch err {
	case nil:
		return c.JSON(createProduct)
	case UserAlreadyExistError, PasswordHashingError:
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}

func (a *Api) GetProducts(c *fiber.Ctx) error {
	productsList, err := a.Service.GetProducts()

	switch err {
	case nil:
		return c.JSON(productsList)
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}
