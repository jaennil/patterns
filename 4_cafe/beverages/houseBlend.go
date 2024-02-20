package beverages

var houseBlendCost float64 = 2

type houseBlend struct {
	cost        float64
	description string
}

func NewHouseBlend() houseBlend {
	return houseBlend{houseBlendCost, "House Blend"}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return houseBlend.description
}
