package flavors

import "3/beverages"

var soyCost float64 = 3

type soy struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewSoy(beverage beverages.IBeverage) soy {
	return soy{beverage, 3, "Soy"}
}

func (flavor soy) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

func (soy soy) Description() string {
	return soy.beverage.Description() + ", " + soy.description
}
