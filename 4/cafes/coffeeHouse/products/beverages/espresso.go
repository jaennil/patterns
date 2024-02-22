package beverages

import "fmt"

type espresso struct {
	beverage
}

func NewEspresso(cost, size float64) espresso {
	return espresso{
		beverage{
			description: "Espresso",
			cost:        cost,
			size:        size,
		},
	}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return fmt.Sprintf("%v %gL", espresso.description, espresso.size)
}
