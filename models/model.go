package models

type RegisterDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Surname  string `json:"surname"`
	Tel      string `json:"tel"`
}

type User struct {
	ID       string `json:"uid"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Password string `json:"-"`
}

type UserEntity struct {
	ID       string `bson:"uid"`
	Name     string `bson:"name"`
	Surname  string `bson:"surname"`
	Email    string `bson:"email"`
	Tel      string `bson:"tel"`
	Password string `bson:"password"`
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
