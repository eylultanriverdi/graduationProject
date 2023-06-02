package main

import (
	"encoding/base64"
	"math"
	"strconv"

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
		CalorieValue:      product.CalorieValue,
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
	page := c.Query("page", "1")
	limit := c.Query("limit", "10") // Varsayılan olarak 10 ürün gösterilecek

	// page ve limit değerlerini integer'a dönüştürme işlemlerini yapabilirsiniz
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid page number")
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid limit value")
	}

	productsList, err := a.Service.GetProducts(pageNum, limitNum)

	switch err {
	case nil:
		// Sayfa sayısını ve toplam ürün sayısını hesapla
		totalProducts, err := a.Service.GetTotalProducts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		totalPages := int(math.Ceil(float64(totalProducts) / float64(limitNum)))

		// Sayfalama bilgilerini dön
		paginationInfo := map[string]interface{}{
			"currentPage":   pageNum,
			"totalPages":    totalPages,
			"totalProducts": totalProducts,
			"perPage":       limitNum,
		}

		response := map[string]interface{}{
			"pagination": paginationInfo,
			"products":   productsList,
		}

		return c.JSON(response)
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}
