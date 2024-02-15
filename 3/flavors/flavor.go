package flavors

import (
	"3/beverages"
	"fmt"
)

var (
	milkCost  float64 = 1
	mochaCost float64 = 2
	soyCost   float64 = 3
	whipCost  float64 = 4
)

func Print() {
	fmt.Println("---Flavors---")
	fmt.Printf("1. Milk %v$\n", milkCost)
	fmt.Printf("2. Mocha %v$\n", mochaCost)
	fmt.Printf("3. Soy %v$\n", soyCost)
	fmt.Printf("4. Whip %v$\n", whipCost)
}

type milk struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewMilk(beverage beverages.IBeverage) milk {
	return milk{beverage, 1, "Milk"}
}

func (milk milk) Description() string {
	return milk.beverage.Description() + ", " + milk.description
}

func (flavor milk) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

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

type soy struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewSoy(beverage beverages.IBeverage) soy {
	return soy{beverage, 3, "Soy"}
}

func (flavor soy) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

func (soy soy) Description() string {
	return soy.beverage.Description() + ", " + soy.description
}

type whip struct {
	beverage    beverages.IBeverage
	cost        float64
	description string
}

func NewWhip(beverage beverages.IBeverage) whip {
	return whip{beverage, 4, "Whip"}
}

func (flavor whip) Cost() float64 {
	return flavor.beverage.Cost() + flavor.cost
}

func (whip whip) Description() string {
	return whip.beverage.Description() + ", " + whip.description
}
