package beverages

type Espresso struct {
	IBeverage
}

func (espresso *Espresso) Cost() int {
	return 3
}

func (espresso *Espresso) GetDescription() string {
	return "Espresso"
}

