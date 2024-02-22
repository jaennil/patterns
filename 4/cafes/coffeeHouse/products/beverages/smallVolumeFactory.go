package beverages

import "fmt"

const (
	Darkroast = iota + 1
	Houseblend
	Decaf
	Espresso
)

type SmallVolumeFactory struct {
}

func (SmallVolumeFactory) CreateBeverage(beverageType int, cost float64) (IBeverage, error) {
	switch beverageType {
	case Darkroast:
		return NewDarkRoast(cost, .4), nil
	case Houseblend:
		return NewHouseBlend(cost, .4), nil
	case Decaf:
		return NewDecaf(cost, .4), nil
	case Espresso:
		return NewEspresso(cost, .4), nil
	default:
		return nil, fmt.Errorf("unknown beverage type")
	}
}
