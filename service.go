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
		ID:                  GenerateUUID(8),
		Name:                register.Name,
		Surname:             register.Surname,
		Email:               register.Email,
		Tel:                 register.Tel,
		Password:            register.Password,
		Age:                 register.Age,
		Kilo:                register.Kilo,
		Height:              register.Height,
		AmountofWater:       register.AmountofWater,
		DailyMovementAmount: register.DailyMovementAmount,
		DesiredWeight:       register.DesiredWeight,
		DesiredDestination:  register.DesiredDestination,
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

func (service *Service) AddRecipe(recipe models.RecipeDTO) (*models.Recipe, error) {
	recipeCreate := models.Recipe{
		RecipeID:     GenerateUUID(8),
		RecipeName:   recipe.RecipeName,
		RecipeDetail: recipe.RecipeDetail,
	}

	addRecipe, err := service.Repository.AddRecipe(recipeCreate)
	if err != nil {
		return nil, err
	}

	return addRecipe, nil
}

func (service *Service) NutritionistRegister(registerNutritionist models.NutritionistRegisterDTO) (*models.Nutritionist, error) {
	nutritionistCreate := models.Nutritionist{
		ID:          GenerateUUID(8),
		Name:        registerNutritionist.Name,
		Surname:     registerNutritionist.Surname,
		Email:       registerNutritionist.Email,
		Tel:         registerNutritionist.Tel,
		Password:    registerNutritionist.Password,
		Age:         registerNutritionist.Age,
		Uni:         registerNutritionist.Uni,
		Experience:  registerNutritionist.Experience,
		Profession:  registerNutritionist.Profession,
		Explanation: registerNutritionist.Explanation,
	}

	_, err := service.Repository.GetByEmail(registerNutritionist.Email)
	if err == nil {
		return nil, UserAlreadyExistError
	}

	createNutritionist, err := service.Repository.CreateNutritionist(nutritionistCreate)
	if err != nil {
		return nil, err
	}

	return createNutritionist, nil
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
		DietitianName:  dietCategory.DietitianName,
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
	createdList, err := service.Repository.CreateCalorieList(calorieList)
	if err != nil {
		return nil, err
	}

	return createdList, nil
}

func (service *Service) CreateNutritionistList(nutritionistList models.NutritionistList) (*models.NutritionistList, error) {
	createdNutritionistList, err := service.Repository.CreateNutritionistList(nutritionistList)
	if err != nil {
		return nil, err
	}

	return createdNutritionistList, nil
}

func (service *Service) GetProducts(page int, limit int) ([]models.Product, error) {
	skip := (page - 1) * limit

	productsList, err := service.Repository.GetProducts(skip, limit)
	if err != nil {
		return nil, err
	}

	return productsList, nil
}

func (service *Service) GetTotalProducts() (int, error) {
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

func (service *Service) GetNutritionists() ([]models.Nutritionist, error) {
	nutritionists, err := service.Repository.GetNutritionists()

	if err != nil {
		return nil, err
	}

	return nutritionists, nil
}

func (service *Service) GetNutritionistList() ([]models.NutritionistList, error) {
	nutritionistList, err := service.Repository.GetNutritionistList()

	if err != nil {
		return nil, err
	}

	return nutritionistList, nil
}

func (service *Service) GetRecipeList() ([]models.Recipe, error) {
	recipeList, err := service.Repository.GetRecipeList()

	if err != nil {
		return nil, err
	}

	return recipeList, nil
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

func (service *Service) SigninNutritionist(signin models.SigninNutritionistDTO) (*models.Nutritionist, error) {
	nutritionist, err := service.Repository.GetByEmailNutritionist(signin.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.UserNotFoundError
		}
		return nil, err
	}

	if nutritionist.Password != signin.Password {
		return nil, models.InvalidPasswordError
	}

	return nutritionist, nil
}

func (service *Service) GetProfile(userID string) (*models.User, error) {
	return service.Repository.GetByID(userID)
}

func (service *Service) GetNutritionistProfile(nutritionistID string) (*models.Nutritionist, error) {
	return service.Repository.GetByNutritionistId(nutritionistID)
}
