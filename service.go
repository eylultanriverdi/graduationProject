package main

import (
	"errors"
	"strings"

	"example.com/greetings/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) Service {
	return Service{
		Repository: Repository,
	}
}

var UserAlreadyExistError error = errors.New("User already exist")
var PasswordHashingError error = errors.New("Error while hashing password")

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}

func (service *Service) Register(register models.RegisterDTO) (*models.User, error) {
	userCreate := models.User{
		ID:       GenerateUUID(8),
		Name:     register.Name,
		Surname:  register.Surname,
		Email:    register.Email,
		Tel:      register.Tel,
		Password: register.Password,
	}

	_, err := service.Repository.GetByEmail(register.Email)
	if err == nil {
		return nil, UserAlreadyExistError
	}

	createUser, err := service.Repository.CreateUser(userCreate)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}

func (service *Service) GetProduct(product models.ProductCategoryDTO) (*models.Product, error) {
	productCreate := models.Product{
		ProductId:         GenerateUUID(8),
		ProductName:       product.ProductName,
		Description:       product.Description,
		ProductImage:      []byte(product.ProductImage),
		ProteinValue:      product.ProteinValue,
		CarbohydrateValue: product.CarbohydrateValue,
		OilValue:          product.OilValue,
		GlutenValue:       product.GlutenValue,
		KetogenicDiet:     product.KetogenicDiet,
		GlutenFree:        product.GlutenFree,
		SaltFree:          product.SaltFree,
		CalorieValue:      product.CalorieValue,
	}

	createProduct, err := service.Repository.CreateProduct(productCreate)

	if err != nil {
		return nil, err
	}

	return createProduct, nil
}

func (service *Service) GetDietCategory(dietCategory models.DietCategoryDTO) (*models.DietCategory, error) {
	dietCategoryCreate := models.DietCategory{
		CategoryId:     GenerateUUID(8),
		CategoryName:   dietCategory.CategoryName,
		Description:    dietCategory.Description,
		CategoryImage:  []byte(dietCategory.CategoryImage),
		AllowedFoods:   dietCategory.AllowedFoods,
		ForbiddenFoods: dietCategory.ForbiddenFoods,
		DailyDietPlan:  dietCategory.DailyDietPlan,
	}

	createProduct, err := service.Repository.CreateDietCategory(dietCategoryCreate)

	if err != nil {
		return nil, err
	}

	return createProduct, nil
}

func (service *Service) CreateCalorieList(calorieList models.CalorieList) (*models.CalorieList, error) {
	// Repository'de kayıt işlemi yap
	createdList, err := service.Repository.CreateCalorieList(calorieList)
	if err != nil {
		return nil, err
	}

	return createdList, nil
}

func (service *Service) GetProducts(page int, limit int) ([]models.Product, error) {
	skip := (page - 1) * limit

	// Sayfalama parametrelerini kullanarak veri getirme işlemini gerçekleştirin
	productsList, err := service.Repository.GetProducts(skip, limit)
	if err != nil {
		return nil, err
	}

	return productsList, nil
}

func (service *Service) GetTotalProducts() (int, error) {
	// Veritabanında toplam ürün sayısını döndüren bir işlev ekleyin
	total, err := service.Repository.GetTotalProducts()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (service *Service) GetDietCategories() ([]models.DietCategory, error) {
	dietCategories, err := service.Repository.GetDietCategories()

	if err != nil {
		return nil, err
	}

	return dietCategories, nil
}

func (service *Service) GetCalorieList() ([]models.CalorieList, error) {
	calorieInfoList, err := service.Repository.GetCalorieList()

	if err != nil {
		return nil, err
	}

	return calorieInfoList, nil
}

var UserNotFoundError error = errors.New("User not found")
var InvalidPasswordError error = errors.New("Invalid password")

func (service *Service) Signin(signin models.SigninDTO) (*models.User, error) {
	user, err := service.Repository.GetByEmail(signin.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.UserNotFoundError
		}
		return nil, err
	}

	if user.Password != signin.Password {
		return nil, models.InvalidPasswordError
	}

	return user, nil
}
func (service *Service) GetProfile(userID string) (*models.User, error) {
	return service.Repository.GetByID(userID)
}
