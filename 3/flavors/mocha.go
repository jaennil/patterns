package flavors

import "3/beverages"

type Mocha struct {
	Beverage beverages.IBeverage
}

func (mocha *Mocha) Cost() int {
	beverageCost := mocha.Beverage.Cost()
	return beverageCost + 2
}

func (mocha *Mocha) GetDescription() string {
	return mocha.Beverage.GetDescription() + " Mocha"
}
