package beverages

type IBeverageFactory interface {
	createBaseIngredient()
	createMainIngredient()
	createTopper()
}

type beverageFactory struct {
}

func (beverageFactory) create() {
	createBaseIngredient()
	createMainIngredient()
	createTopper()
}
