package models

import "errors"

var (
	UserNotFoundError     = errors.New("User not found")
	InvalidPasswordError  = errors.New("Invalid password")
	UserAlreadyExistError = errors.New("User already exists")
	PasswordHashingError  = errors.New("Error occurred during password hashing")
)

type RegisterDTO struct {
	ID                  string `json:"uid"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	Email               string `json:"email"`
	Tel                 string `json:"tel"`
	Password            string `json:"password"`
	Age                 string `json:"age"`
	Kilo                string `json:"kilo"`
	Height              string `json:"height"`
	AmountofWater       string `json:"amountOfWater"`
	DailyMovementAmount string `json:"dailyMovementAmount"`
	DesiredWeight       string `json:"desiredWeight"`
	DesiredDestination  string `json:"desiredDestination"`
}

type NutritionistRegisterDTO struct {
	ID          string `json:"uid"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Tel         string `json:"tel"`
	Password    string `json:"password"`
	Age         string `json:"age"`
	Uni         string `json:"uni"`
	Experience  string `json:"experience"`
	Profession  string `json:"profession"`
	Explanation string `json:"explanation"`
}

type Nutritionist struct {
	ID          string `json:"uid"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Tel         string `json:"tel"`
	Password    string `json:"password"`
	Age         string `json:"age"`
	Uni         string `json:"uni"`
	Experience  string `json:"experience"`
	Profession  string `json:"profession"`
	Explanation string `json:"explanation"`
}

type NutritionistEntity struct {
	ID          string `bson:"uid"`
	Name        string `bson:"name"`
	Surname     string `bson:"surname"`
	Email       string `bson:"email"`
	Tel         string `bson:"tel"`
	Password    string `bson:"password"`
	Age         string `bson:"age"`
	Uni         string `bson:"uni"`
	Experience  string `bson:"experience"`
	Profession  string `bson:"profession"`
	Explanation string `bson:"explanation"`
}

type User struct {
	ID                  string `json:"uid"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	Email               string `json:"email"`
	Tel                 string `json:"tel"`
	Password            string `json:"password"`
	Age                 string `json:"age"`
	Kilo                string `json:"kilo"`
	Height              string `json:"height"`
	AmountofWater       string `json:"amountOfWater"`
	DailyMovementAmount string `json:"dailyMovementAmount"`
	DesiredWeight       string `json:"desiredWeight"`
	DesiredDestination  string `json:"desiredDestination"`
}

type UserEntity struct {
	ID                  string `bson:"uid"`
	Name                string `bson:"name"`
	Surname             string `bson:"surname"`
	Email               string `bson:"email"`
	Tel                 string `bson:"tel"`
	Password            string `bson:"password"`
	Age                 string `bson:"age"`
	Kilo                string `bson:"kilo"`
	Height              string `bson:"height"`
	AmountofWater       string `bson:"amountOfWater"`
	DailyMovementAmount string `bson:"dailyMovementAmount"`
	DesiredWeight       string `bson:"desiredWeight"`
	DesiredDestination  string `bson:"desiredDestination"`
}

type Product struct {
	ProductId         string `json:"productId"`
	ProductName       string `json:"productName"`
	Description       string `json:"description"`
	ProductImage      []byte `json:"productImage"`
	ProteinValue      string `json:"proteinValue"`
	CarbohydrateValue string `json:"carbohydrateValue"`
	OilValue          string `json:"oilValue"`
	GlutenValue       string `json:"glutenValue"`
	KetogenicDiet     string `json:"ketogenicDiet"`
	GlutenFree        string `json:"glutenFree"`
	SaltFree          string `json:"saltFree"`
	CalorieValue      string `json:"calorieValue"`
}

type ProductEntity struct {
	ProductId         string `bson:"productId"`
	ProductName       string `bson:"productName"`
	Description       string `bson:"description"`
	ProductImage      []byte `bson:"productImage"`
	ProteinValue      string `json:"proteinValue"`
	CarbohydrateValue string `json:"carbohydrateValue"`
	OilValue          string `json:"oilValue"`
	GlutenValue       string `json:"glutenValue"`
	KetogenicDiet     string `json:"ketogenicDiet"`
	GlutenFree        string `json:"glutenFree"`
	SaltFree          string `json:"saltFree"`
	CalorieValue      string `json:"calorieValue"`
}

type ProductCategoryDTO struct {
	ProductId         string `json:"productId"`
	ProductName       string `json:"productName"`
	Description       string `json:"description"`
	ProductImage      string `json:"productImage"`
	ProteinValue      string `json:"proteinValue"`
	CarbohydrateValue string `json:"carbohydrateValue"`
	OilValue          string `json:"oilValue"`
	GlutenValue       string `json:"glutenValue"`
	KetogenicDiet     string `json:"ketogenicDiet"`
	GlutenFree        string `json:"glutenFree"`
	SaltFree          string `json:"saltFree"`
	CalorieValue      string `json:"calorieValue"`
}

type CalorieList struct {
	CalorieListId string    `json:"calorieListId"`
	Products      []Product `json:"products"`
	TotalCalorie  string    `json:"totalCalorie"`
	CreateDate    string    `json:"createDate"`
}

type DietCategoryEntity struct {
	CategoryId     string           `bson:"categoryId"`
	CategoryName   string           `bson:"categoryName"`
	Description    string           `bson:"description"`
	CategoryImage  []byte           `bson:"categoryImage"`
	AllowedFoods   []AllowedFoods   `bson:"allowedFoods"`
	ForbiddenFoods []ForbiddenFoods `bson:"forbiddenFoods"`
	DailyDietPlan  []DailyDietPlan  `bson:"dailyDietPlan"`
}

type DietCategoryDTO struct {
	CategoryId     string           `json:"categoryId"`
	CategoryName   string           `json:"categoryName"`
	Description    string           `json:"description"`
	CategoryImage  []byte           `json:"categoryImage"`
	AllowedFoods   []AllowedFoods   `json:"allowedFoods"`
	ForbiddenFoods []ForbiddenFoods `json:"forbiddenFoods"`
	DailyDietPlan  []DailyDietPlan  `json:"dailyDietPlan"`
}

type DietCategory struct {
	CategoryId     string           `json:"categoryId"`
	CategoryName   string           `json:"categoryName"`
	Description    string           `json:"description"`
	CategoryImage  []byte           `json:"categoryImage"`
	AllowedFoods   []AllowedFoods   `json:"allowedFoods"`
	ForbiddenFoods []ForbiddenFoods `json:"forbiddenFoods"`
	DailyDietPlan  []DailyDietPlan  `json:"dailyDietPlan"`
}

type AllowedFoods struct {
	ProductId         string `json:"productId"`
	ProductName       string `json:"productName"`
	Description       string `json:"description"`
	ProductImage      []byte `json:"productImage"`
	ProteinValue      string `json:"proteinValue"`
	CarbohydrateValue string `json:"carbohydrateValue"`
	OilValue          string `json:"oilValue"`
	GlutenValue       string `json:"glutenValue"`
	KetogenicDiet     string `json:"ketogenicDiet"`
	GlutenFree        string `json:"glutenFree"`
	SaltFree          string `json:"saltFree"`
	CalorieValue      string `json:"calorieValue"`
}

type ForbiddenFoods struct {
	ProductId         string `json:"productId"`
	ProductName       string `json:"productName"`
	Description       string `json:"description"`
	ProductImage      []byte `json:"productImage"`
	ProteinValue      string `json:"proteinValue"`
	CarbohydrateValue string `json:"carbohydrateValue"`
	OilValue          string `json:"oilValue"`
	GlutenValue       string `json:"glutenValue"`
	KetogenicDiet     string `json:"ketogenicDiet"`
	GlutenFree        string `json:"glutenFree"`
	SaltFree          string `json:"saltFree"`
	CalorieValue      string `json:"calorieValue"`
}

type DailyDietPlan struct {
	PlanId         string           `json:"planId"`
	AgeRange       string           `json:"ageRange"`
	ProgramDetails []ProgramDetails `json:"programDetails"`
}

type ProgramDetails struct {
	Breakfast     string `json:"breakfast"`
	Lunch         string `json:"lunch"`
	Dinner        string `json:"dinner"`
	Snack         string `json:"snack"`
	AmountofWater string `json:"amountofWater"`
}

type SigninDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninNutritionistDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
