package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"example.com/greetings/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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
	fmt.Println(createUser, "createUser")
	if err != nil {
		switch err {
		case UserAlreadyExistError, PasswordHashingError:
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		default:
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	return c.JSON(createUser)
}

func (api *Api) NutritionistRegisterHandler(c *fiber.Ctx) error {
	nutritionistRegister := models.NutritionistRegisterDTO{}
	err := c.BodyParser(&nutritionistRegister)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	createNutritionist, err := api.Service.NutritionistRegister(nutritionistRegister)
	if err != nil {
		switch err {
		case UserAlreadyExistError, PasswordHashingError:
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		default:
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	return c.JSON(createNutritionist)
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

func (api *Api) SigninHandler(c *fiber.Ctx) error {
	signin := models.SigninDTO{}
	err := c.BodyParser(&signin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Kullanıcının kimlik doğrulaması gerçekleştiriliyor
	user, err := api.Service.Signin(signin)
	if err != nil {
		if err == models.UserNotFoundError {
			return c.Status(fiber.StatusUnauthorized).SendString("User not found")
		} else if err == models.InvalidPasswordError {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	// Token oluşturuluyor
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID

	// Token imzalanıyor
	tokenString, err := token.SignedString([]byte("xL#j9E7o!P1k@9qR3tZw5y"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Token yanıt olarak döndürülüyor
	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func (api *Api) SigninNutritionistHandler(c *fiber.Ctx) error {
	signin := models.SigninNutritionistDTO{}
	err := c.BodyParser(&signin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Kullanıcının kimlik doğrulaması gerçekleştiriliyor
	nutritionist, err := api.Service.SigninNutritionist(signin)
	if err != nil {
		if err == models.UserNotFoundError {
			return c.Status(fiber.StatusUnauthorized).SendString("User not found")
		} else if err == models.InvalidPasswordError {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	// Token oluşturuluyor
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = nutritionist.ID

	// Token imzalanıyor
	tokenString, err := token.SignedString([]byte("xL#j9E7o!P1k@9qR3tZw5y"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Token yanıt olarak döndürülüyor
	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func (api *Api) ProfileHandler(c *fiber.Ctx) error {
	tokenString := strings.TrimSpace(strings.TrimPrefix(c.Get("Authorization"), "Bearer "))

	// Tokenı isteğin başlığından alınabilir

	// Token doğrulanıyor
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("xL#j9E7o!P1k@9qR3tZw5y"), nil
	})

	fmt.Println(token, "token")

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Profil bilgileri alınıyor
	user, err := api.Service.GetProfile(userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(user)
}

func (api *Api) ProfileNutritionistHandler(c *fiber.Ctx) error {
	tokenString := strings.TrimSpace(strings.TrimPrefix(c.Get("Authorization"), "Bearer "))

	// Tokenı isteğin başlığından alınabilir

	// Token doğrulanıyor
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("xL#j9E7o!P1k@9qR3tZw5y"), nil
	})

	fmt.Println(token, "token")

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	nutritionistId, ok := claims["nutritionistID"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Profil bilgileri alınıyor
	nutritionist, err := api.Service.GetNutritionistProfile(nutritionistId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).SendString("Nutritionist not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(nutritionist)
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

func (a *Api) GetDietCategories(c *fiber.Ctx) error {
	calorieInfoList, err := a.Service.GetDietCategories()

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	c.JSON(calorieInfoList)
	return nil
}

func (a *Api) GetCalorieList(c *fiber.Ctx) error {
	calorieInfoList, err := a.Service.GetCalorieList()

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return err
	}

	c.JSON(calorieInfoList)
	return nil
}
