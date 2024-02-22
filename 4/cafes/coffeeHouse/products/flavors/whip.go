package flavors

import "3/cafes/coffeeHouse/products/beverages"

type whip struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewWhip(beverage beverages.IBeverage) whip {
	return whip{beverage, 4, "Whip"}
}

func (flavor whip) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

func (whip whip) Description() string {
	return whip.beverage.Description() + ", " + whip.description
}
