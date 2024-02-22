package beverages

import "fmt"

type MediumVolumeFactory struct {
}

func (MediumVolumeFactory) CreateBeverage(beverageType int, cost float64) (IBeverage, error) {
	switch beverageType {
	case Darkroast:
		return NewDarkRoast(cost, .6), nil
	case Houseblend:
		return NewHouseBlend(cost, .6), nil
	case Decaf:
		return NewDecaf(cost, .6), nil
	case Espresso:
		return NewEspresso(cost, .6), nil
	default:
		return nil, fmt.Errorf("unknown beverage type")
	}
}
