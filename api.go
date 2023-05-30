package main

import (
	"encoding/base64"
	"net/http"

	"example.com/greetings/models"
	"github.com/gofiber/fiber"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) RegisterHandler(c *fiber.Ctx) {
	register := models.RegisterDTO{}
	err := c.BodyParser(&register)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return
	}
	createUser, err := api.Service.Register(register)

	switch err {
	case nil:
		c.JSON(createUser)
		c.Status(http.StatusCreated)
	case UserAlreadyExistError, PasswordHashingError:
		c.Status(fiber.StatusBadRequest)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (api *Api) ProductHandler(c *fiber.Ctx) {
	product := models.ProductCategoryDTO{}
	err := c.BodyParser(&product)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return
	}

	// Decode base64 string to byte array
	productImage, err := base64.StdEncoding.DecodeString(product.ProductImage)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return
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
		c.JSON(createProduct)
		c.Status(http.StatusCreated)
	case UserAlreadyExistError, PasswordHashingError:
		c.Status(fiber.StatusBadRequest)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func (a *Api) GetProducts(c *fiber.Ctx) {

	productsList, err := a.Service.GetProducts()

	switch err {
	case nil:
		c.JSON(productsList)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}
