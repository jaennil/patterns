package flavors

import "3/beverages"

type Soy struct {
	Beverage beverages.IBeverage
}

func (soy *Soy) Cost() int {
	beverageCost := soy.Beverage.Cost()
	return beverageCost + 3
}

func (soy *Soy) GetDescription() string {
	return soy.Beverage.GetDescription() + " soy"
}

