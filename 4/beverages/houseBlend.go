package beverages

var houseBlendCost float64 = 2

type houseBlend struct {
	beverage
}

func NewHouseBlend(size float64) houseBlend {
	return houseBlend{
		beverage{
			cost:        houseBlendCost,
			description: "House Blend",
			size:        size,
		},
	}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return houseBlend.description
}
