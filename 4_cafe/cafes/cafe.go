package cafes

import (
	"3/beverages"
	"3/flavors"
	"fmt"
)

type beverageType int
type flavorType int

const (
	milk flavorType = iota + 1
	mocha
	soy
	whip
)

const (
	darkroast beverageType = iota + 1
	houseblend
	decaf
	espresso
)

type cafe struct {
	order []beverages.IBeverage
}

func NewCafe() cafe {
	return cafe{}
}

func PrintMenu() {
	fmt.Println("*****Menu*****")
	beverages.Print()
	flavors.Print()
	fmt.Println("**************")
}

func (cafe) PrintOrderCost() {
	fmt.Println("---Order---")
	// fmt.Println(beverage.Description())
	// fmt.Println(beverage.Cost(), "$")
	fmt.Println("-----------")
	fmt.Println()
}


func (cafe cafe) newOrder() {
	cafe.order = nil
}

func (cafe cafe) OrderBeverage(beverageType beverageType) error {

	switch beverageType {
	case darkroast:
		cafe.order = append(cafe.order, beverages.NewDarkRoast())
		return nil
	case houseblend:
		cafe.order = append(cafe.order,beverages.NewHouseBlend())
		return nil
	case decaf:
		cafe.order = append(cafe.order,beverages.NewDecaf())
		return nil
	case espresso:
		cafe.order = append(cafe.order,beverages.NewEspresso())
		return nil
	}

	return fmt.Errorf("unknown beverage type")
}

func (cafe) AddFlavor(beverage beverages.IBeverage, flavorType flavorType) (beverages.IBeverage, error) {
	switch flavorType {

	case milk:
		return flavors.NewMilk(beverage), nil
	case mocha:
		return flavors.NewMocha(beverage), nil
	case soy:
		return flavors.NewSoy(beverage), nil
	case whip:
		return flavors.NewWhip(beverage), nil
	}

	return nil, fmt.Errorf("unknown flavor type")
}
