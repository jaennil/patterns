package flavors

import (
	"3/beverages"
	"fmt"
)

var (
	milkCost float64 = 1
	mochaCost float64 = 2
	soyCost float64 = 3
	whipCost float64 = 4
)

func Print() {
	fmt.Println("---Flavors---")
	fmt.Printf("1. Milk %v$\n", milkCost)
	fmt.Printf("2. Mocha %v$\n", mochaCost)
	fmt.Printf("3. Soy %v$\n", soyCost)
	fmt.Printf("4. Whip %v$\n", whipCost)
}

type milk struct {
	Beverage beverages.IBeverage
	cost     float64
}

func NewMilk(beverage beverages.IBeverage) milk {
	return milk{beverage, 1}
}

func (flavor milk) Cost() float64 {
	return flavor.Beverage.Cost() + flavor.cost
}

type mocha struct {
	Beverage beverages.IBeverage
	cost     float64
}

func NewMocha(beverage beverages.IBeverage) mocha {
	return mocha{beverage, 2}
}

func (flavor mocha) Cost() float64 {
	return flavor.Beverage.Cost() + flavor.cost
}

type soy struct {
	Beverage beverages.IBeverage
	cost     float64
}

func NewSoy(beverage beverages.IBeverage) soy {
	return soy{beverage, 3}
}

func (flavor soy) Cost() float64 {
	return flavor.Beverage.Cost() + flavor.cost
}

type whip struct {
	Beverage beverages.IBeverage
	cost     float64
}

func NewWhip(beverage beverages.IBeverage) whip {
	return whip{beverage, 4}
}

func (flavor whip) Cost() float64 {
	return flavor.Beverage.Cost() + flavor.cost
}
