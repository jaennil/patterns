package beverages

type houseBlend struct {
	description string
	cost        float64
}

func NewHouseBlend(cost float64) houseBlend {
	return houseBlend{
		description: "House Blend",
		cost:        cost,
	}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return houseBlend.description
}
