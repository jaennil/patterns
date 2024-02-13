package flavors

import "3/beverages"

type Milk struct {
	Beverage beverages.IBeverage
}

func (milk *Milk) Cost() int {
	beverageCost := milk.Beverage.Cost()
	return beverageCost + 1
}

func (milk *Milk) GetDescription() string {
	return milk.Beverage.GetDescription() + " milk"
}

