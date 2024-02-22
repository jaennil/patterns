package flavors

import "3/cafes/coffeeHouse/products/beverages"

type mocha struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewMocha(beverage beverages.IBeverage) mocha {
	return mocha{beverage, 2, "Mocha"}
}

func (flavor mocha) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

func (mocha mocha) Description() string {
	return mocha.beverage.Description() + ", " + mocha.description
}
