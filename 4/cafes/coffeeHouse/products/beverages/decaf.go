package beverages

import "fmt"

type decaf struct {
	beverage
}

func NewDecaf(cost, size float64) decaf {
	return decaf{
		beverage{
			description: "Decaf",
			cost:        cost,
			size:        size,
		},
	}
}

func (decaf decaf) Cost() float64 {
	return decaf.cost
}

func (decaf decaf) Description() string {
	return fmt.Sprintf("%v %gL", decaf.description, decaf.size)
}
