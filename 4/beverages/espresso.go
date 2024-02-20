package beverages

var espressoCost float64 = 3

type espresso struct {
	beverage
}

func NewEspresso(size float64) espresso {
	return espresso{
		beverage{
			cost:        espressoCost,
			description: "Espresso",
			size:        size,
		},
	}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return espresso.description
}
