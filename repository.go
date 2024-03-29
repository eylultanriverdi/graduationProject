package main

import (
	"context"
	"log"
	"time"

	"example.com/greetings/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() *Repository {
	uri := "mongodb+srv://project-Test:projectTest@cluster0.t0wnmxh.mongodb.net/test"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client: client}
}

func NewTestRepository() *Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client: client}
}

func GetCleanTestRepository() *Repository {
	repository := NewTestRepository()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	blogDB := repository.client.Database("user")
	err := blogDB.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return repository
}

func ConvertUserEntityToUser(userEntity models.UserEntity) models.User {
	return models.User{
		ID:                  userEntity.ID,
		Name:                userEntity.Name,
		Surname:             userEntity.Surname,
		Email:               userEntity.Email,
		Tel:                 userEntity.Tel,
		Password:            userEntity.Password,
		Age:                 userEntity.Age,
		Kilo:                userEntity.Kilo,
		Height:              userEntity.Height,
		AmountofWater:       userEntity.AmountofWater,
		DailyMovementAmount: userEntity.DailyMovementAmount,
		DesiredWeight:       userEntity.DesiredWeight,
		DesiredDestination:  userEntity.DesiredDestination,
	}
}

func ConvertRecipeEntityToNutritionist(recipeEntity models.RecipeEntity) models.Recipe {
	return models.Recipe{
		RecipeID:     recipeEntity.RecipeID,
		RecipeName:   recipeEntity.RecipeName,
		RecipeDetail: recipeEntity.RecipeDetail,
	}
}

func ConvertNutritionistEntityToNutritionist(nutritionistEntity models.NutritionistEntity) models.Nutritionist {
	return models.Nutritionist{
		ID:          nutritionistEntity.ID,
		Name:        nutritionistEntity.Name,
		Surname:     nutritionistEntity.Surname,
		Email:       nutritionistEntity.Email,
		Tel:         nutritionistEntity.Tel,
		Password:    nutritionistEntity.Password,
		Age:         nutritionistEntity.Age,
		Uni:         nutritionistEntity.Uni,
		Experience:  nutritionistEntity.Experience,
		Profession:  nutritionistEntity.Profession,
		Explanation: nutritionistEntity.Explanation,
	}
}

func ConvertProductEntityToProduct(productEntity models.ProductEntity) models.Product {
	return models.Product{
		ProductId:         productEntity.ProductId,
		ProductName:       productEntity.ProductName,
		Description:       productEntity.Description,
		ProductImage:      productEntity.ProductImage,
		ProteinValue:      productEntity.ProteinValue,
		CarbohydrateValue: productEntity.CarbohydrateValue,
		OilValue:          productEntity.OilValue,
		GlutenValue:       productEntity.GlutenValue,
		KetogenicDiet:     productEntity.KetogenicDiet,
		GlutenFree:        productEntity.GlutenFree,
		SaltFree:          productEntity.SaltFree,
		CalorieValue:      productEntity.CalorieValue,
	}
}

func ConvertDietCategoryEntityToCategory(dietCategoryEntity models.DietCategoryEntity) models.DietCategory {
	return models.DietCategory{
		CategoryId:     dietCategoryEntity.CategoryId,
		CategoryName:   dietCategoryEntity.CategoryName,
		DietitianName:  dietCategoryEntity.DietitianName,
		Description:    dietCategoryEntity.Description,
		CategoryImage:  dietCategoryEntity.CategoryImage,
		AllowedFoods:   dietCategoryEntity.AllowedFoods,
		ForbiddenFoods: dietCategoryEntity.ForbiddenFoods,
		DailyDietPlan:  dietCategoryEntity.DailyDietPlan,
	}
}

func (repository *Repository) CreateUser(user models.User) (*models.User, error) {
	collection := repository.client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.UserEntity(user)
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return repository.GetByUserId(user.ID)
}

func (repository *Repository) AddRecipe(recipe models.Recipe) (*models.Recipe, error) {
	collection := repository.client.Database("recipes").Collection("recipe")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.RecipeEntity(recipe)
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return repository.GetByRecipeId(recipe.RecipeID)
}

func (repository *Repository) CreateNutritionist(nutritionist models.Nutritionist) (*models.Nutritionist, error) {
	collection := repository.client.Database("nutritionists").Collection("nutritionist")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.NutritionistEntity(nutritionist)
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return repository.GetByNutritionistId(nutritionist.ID)
}

func (repository *Repository) CreateProduct(product models.Product) (*models.Product, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.ProductEntity(product)
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return repository.GetByProductId(product.ProductId)
}

func (repository *Repository) CreateDietCategory(dietCategory models.DietCategory) (*models.DietCategory, error) {
	collection := repository.client.Database("categories").Collection("category")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.DietCategoryEntity(dietCategory)
	_, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return repository.GetByCategoryId(dietCategory.CategoryId)
}

func (repository *Repository) CreateCalorieList(calorieList models.CalorieList) (*models.CalorieList, error) {
	collection := repository.client.Database("calorieLists").Collection("calorieList")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, calorieList)
	if err != nil {
		return nil, err
	}

	return &calorieList, nil
}

func (repository *Repository) CreateNutritionistList(nutritionistList models.NutritionistList) (*models.NutritionistList, error) {
	collection := repository.client.Database("nutritionistsList").Collection("nutritionistList")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, nutritionistList)
	if err != nil {
		return nil, err
	}

	return &nutritionistList, nil
}
func (repository *Repository) GetProducts(skip int, limit int) ([]models.Product, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"_id": -1})
	options.SetSkip(int64(skip))
	options.SetLimit(int64(limit))

	cur, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	entities := []models.Product{}

	for cur.Next(ctx) {
		entity := models.ProductEntity{}
		err := cur.Decode(&entity)
		if err != nil {
			log.Fatal(err)
		}
		product := ConvertProductEntityToProduct(entity)
		entities = append(entities, product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository *Repository) GetByUserId(userId string) (*models.User, error) {
	collection := repository.client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.UserEntity{}

	filter := bson.M{"uid": userId}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	user := ConvertUserEntityToUser(entity)
	return &user, nil
}

func (repository *Repository) GetByRecipeId(recipeID string) (*models.Recipe, error) {
	collection := repository.client.Database("recipes").Collection("recipe")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.RecipeEntity{}

	filter := bson.M{"recipeID": recipeID}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	recipe := ConvertRecipeEntityToNutritionist(entity)
	return &recipe, nil
}

func (repository *Repository) GetByNutritionistId(nutritionistId string) (*models.Nutritionist, error) {
	collection := repository.client.Database("nutritionists").Collection("nutritionist")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.NutritionistEntity{}

	filter := bson.M{"uid": nutritionistId}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	nutritionist := ConvertNutritionistEntityToNutritionist(entity)
	return &nutritionist, nil
}

func (repository *Repository) GetByProductId(productId string) (*models.Product, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.ProductEntity{}

	filter := bson.M{"productId": productId}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	product := ConvertProductEntityToProduct(entity)
	return &product, nil
}

func (repository *Repository) GetByCategoryId(categoryId string) (*models.DietCategory, error) {
	collection := repository.client.Database("categories").Collection("category")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.DietCategoryEntity{}

	filter := bson.M{"categoryId": categoryId}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	dietCategory := ConvertDietCategoryEntityToCategory(entity)
	return &dietCategory, nil
}

func (repository *Repository) GetTotalProducts() (int, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (repository *Repository) GetDietCategories() ([]models.DietCategory, error) {
	collection := repository.client.Database("categories").Collection("category")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dietCategory []models.DietCategory

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dietCategories models.DietCategory
		if err := cursor.Decode(&dietCategories); err != nil {
			return nil, err
		}
		dietCategory = append(dietCategory, dietCategories)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dietCategory, nil
}

func (repository *Repository) GetNutritionists() ([]models.Nutritionist, error) {
	collection := repository.client.Database("nutritionists").Collection("nutritionist")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var nutritionists []models.Nutritionist

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var nutritionistsList models.Nutritionist
		if err := cursor.Decode(&nutritionistsList); err != nil {
			return nil, err
		}
		nutritionists = append(nutritionists, nutritionistsList)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nutritionists, nil
}

func (repository *Repository) GetNutritionistList() ([]models.NutritionistList, error) {
	collection := repository.client.Database("nutritionistsList").Collection("nutritionistList")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var nutritionistList []models.NutritionistList

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var nutritionists models.NutritionistList
		if err := cursor.Decode(&nutritionists); err != nil {
			return nil, err
		}
		nutritionistList = append(nutritionistList, nutritionists)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nutritionistList, nil
}

func (repository *Repository) GetRecipeList() ([]models.Recipe, error) {
	collection := repository.client.Database("recipes").Collection("recipe")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var recipeList []models.Recipe

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var recipes models.Recipe
		if err := cursor.Decode(&recipes); err != nil {
			return nil, err
		}
		recipeList = append(recipeList, recipes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return recipeList, nil
}

func (repository *Repository) GetCalorieList() ([]models.CalorieList, error) {
	collection := repository.client.Database("calorieLists").Collection("calorieList")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var calorieInfoList []models.CalorieList

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var calorieInfo models.CalorieList
		if err := cursor.Decode(&calorieInfo); err != nil {
			return nil, err
		}
		calorieInfoList = append(calorieInfoList, calorieInfo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return calorieInfoList, nil
}

func (repository *Repository) GetByEmail(email string) (*models.User, error) {
	collection := repository.client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.UserEntity{}

	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	user := ConvertUserEntityToUser(entity)

	return &user, nil
}

func (repository *Repository) GetByEmailNutritionist(email string) (*models.Nutritionist, error) {
	collection := repository.client.Database("nutritionists").Collection("nutritionist")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := models.NutritionistEntity{}

	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}

	nutritionist := ConvertNutritionistEntityToNutritionist(entity)

	return &nutritionist, nil
}

func (repository *Repository) GetByID(userID string) (*models.User, error) {
	collection := repository.client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"uid": userID}

	entity := models.UserEntity{}
	err := collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}

	user := models.User{
		ID:                  entity.ID,
		Name:                entity.Name,
		Surname:             entity.Surname,
		Email:               entity.Email,
		Tel:                 entity.Tel,
		Password:            entity.Password,
		Age:                 entity.Age,
		Kilo:                entity.Kilo,
		Height:              entity.Height,
		AmountofWater:       entity.AmountofWater,
		DailyMovementAmount: entity.DailyMovementAmount,
		DesiredWeight:       entity.DesiredWeight,
		DesiredDestination:  entity.DesiredDestination,
	}

	return &user, nil
}
