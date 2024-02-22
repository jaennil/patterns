package beverages

import "fmt"

type LargeVolumeFactory struct {
}

func (LargeVolumeFactory) CreateBeverage(beverageType int, cost float64) (IBeverage, error) {
	switch beverageType {
	case Darkroast:
		return NewDarkRoast(cost, .8), nil
	case Houseblend:
		return NewHouseBlend(cost, .8), nil
	case Decaf:
		return NewDecaf(cost, .8), nil
	case Espresso:
		return NewEspresso(cost, .8), nil
	default:
		return nil, fmt.Errorf("unknown beverage type")
	}
}
