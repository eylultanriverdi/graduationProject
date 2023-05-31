package main

import (
	"errors"
	"strings"

	"example.com/greetings/models"
	"github.com/google/uuid"
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

	if err != nil {
		return nil, err
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

func (service *Service) GetProducts() ([]models.Product, error) {
	productsList, err := service.Repository.GetProducts()

	if err != nil {
		return nil, err
	}

	return productsList, nil
}
