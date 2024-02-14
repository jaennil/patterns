package flavors

import "3/beverages"

type Whip struct {
	Beverage beverages.IBeverage
}

func (whip *Whip) Cost() int {
	beverageCost := whip.Beverage.Cost()
	return beverageCost + 4
}

func (whip *Whip) GetDescription() string {
	return whip.Beverage.GetDescription() + " Whip"
}
