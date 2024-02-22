package beverages

import "fmt"

type houseBlend struct {
	beverage
}

func NewHouseBlend(cost, size float64) houseBlend {
	return houseBlend{
		beverage{
			description: "House Blend",
			cost:        cost,
			size:        size,
		},
	}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return fmt.Sprintf("%v %gL", houseBlend.description, houseBlend.size)
}
