package beverages

type HouseBlend struct {
	IBeverage
}

func (houseBlend *HouseBlend) Cost() int {
	return 4
}

func (houseBlend *HouseBlend) GetDescription() string {
	return "houseBlend"
}


