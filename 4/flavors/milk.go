package flavors

import "3/beverages"

var milkCost float64 = 1

type milk struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewMilk(beverage beverages.IBeverage) milk {
	return milk{beverage, 1, "Milk"}
}

func (milk milk) Description() string {
	return milk.beverage.Description() + ", " + milk.description
}

func (flavor milk) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}
