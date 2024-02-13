package factories

type BeverageFactory interface {
	CreateBase() Base
	CreateMainIngredient() MainIngredient
	CreateTopper() Topper
}