package main

import (
	"encoding/base64"
	"math"
	"strconv"
	"time"

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

func (api *Api) DietCategoryHandler(c *fiber.Ctx) error {
	dietCategory := models.DietCategoryDTO{}
	err := c.BodyParser(&dietCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Create Product instance with byte array
	createCategory, err := api.Service.GetDietCategory(models.DietCategoryDTO{
		CategoryId:     dietCategory.CategoryId,
		CategoryName:   dietCategory.CategoryName,
		Description:    dietCategory.Description,
		CategoryImage:  []byte(dietCategory.CategoryImage),
		AllowedFoods:   dietCategory.AllowedFoods,
		ForbiddenFoods: dietCategory.ForbiddenFoods,
		DailyDietPlan:  dietCategory.DailyDietPlan,
	})

	switch err {
	case nil:
		return c.JSON(createCategory)
	case UserAlreadyExistError, PasswordHashingError:
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	default:
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}

func (api *Api) HandleAddListProduct(c *fiber.Ctx) error {
	calorieList := models.CalorieList{}
	err := c.BodyParser(&calorieList)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// TotalCalorie hesaplaması
	totalCalorie := 0
	for _, product := range calorieList.Products {
		calorie, err := strconv.Atoi(product.CalorieValue)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		totalCalorie += calorie
	}
	calorieList.TotalCalorie = strconv.Itoa(totalCalorie)

	currentTime := time.Now().Format("2006-01-02")
	calorieList.CreateDate = currentTime

	// CalorieListId oluşturulması
	calorieList.CalorieListId = GenerateUUID(8)

	// Repository'de kayıt işlemi yap
	createdList, err := api.Service.CreateCalorieList(calorieList)
	if err != nil {
		switch err {
		case UserAlreadyExistError, PasswordHashingError:
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		default:
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	return c.JSON(createdList)
}

func (a *Api) GetProducts(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "12") // Varsayılan olarak 10 ürün gösterilecek

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

func (a *Api) GetCalorieList(c *fiber.Ctx) error {
	calorieInfoList, err := a.Service.GetKartelamDiskList()

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	c.JSON(calorieInfoList)
	return nil
}
