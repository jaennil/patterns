package beverages

type houseBlend struct {
	description string
	cost        float64
	size        float64
}

func NewHouseBlend(cost, size float64) houseBlend {
	return houseBlend{
		description: "House Blend",
		cost:        cost,
		size:        size,
	}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return houseBlend.description
}
