package cafes

import (
	"3/cafes/coffeeHouse/products/beverages"
	"3/cafes/coffeeHouse/products/flavors"
	"fmt"
)

var (
	darkRoastCost  float64 = 1
	houseBlendCost float64 = 2
	espressoCost   float64 = 3
	decafCost      float64 = 4
)

var (
	milkCost  float64 = 1
	mochaCost float64 = 2
	soyCost   float64 = 3
	whipCost  float64 = 4
)

const (
	milk = iota + 1
	mocha
	soy
	whip
)

var beveragesMenu = fmt.Sprintf(`
---Beverages---
1. Dark Roast %v$
2. House Blend %v$
3. Decaf %v$
4. Espresso %v$
---------------
`, darkRoastCost, houseBlendCost, decafCost, espressoCost)

var flavorsMenu = fmt.Sprintf(`
---Flavors---
1. Milk %v$
2. Mocha %v$
3. Soy %v$
4. Whip %v$
-------------
`, milkCost, mochaCost, soyCost, whipCost)

type beverageCafe struct {
	Cafe
}

func NewBeverageCafe() beverageCafe {
	return beverageCafe{}
}

func (beverageCafe) PrintOperations() {
	fmt.Println("*****Operations*****")
	fmt.Println("1. Print menu")
	fmt.Println("2. Order beverage")
	fmt.Println("3. Add flavor")
	fmt.Println("4. New Order")
	fmt.Println("any. Exit")
	fmt.Println("********************")
}

func (cafe beverageCafe) BeveragesMenu() string {
	return beveragesMenu
}

func (cafe beverageCafe) FlavorsMenu() string {
	return flavorsMenu
}

func (cafe beverageCafe) PrintMenu() {
	fmt.Println("*****Menu*****")
	fmt.Println("---Beverages---")
	fmt.Printf("1. Dark Roast %v$\n", darkRoastCost)
	fmt.Printf("2. House Blend %v$\n", houseBlendCost)
	fmt.Printf("3. Decaf %v$\n", decafCost)
	fmt.Printf("4. Espresso %v$\n", espressoCost)
	fmt.Println("---Flavors---")
	fmt.Printf("1. Milk %v$\n", milkCost)
	fmt.Printf("2. Mocha %v$\n", mochaCost)
	fmt.Printf("3. Soy %v$\n", soyCost)
	fmt.Printf("4. Whip %v$\n", whipCost)
	fmt.Println("**************")
}

func (cafe beverageCafe) AddFlavor(flavorType int) error {

	if len(cafe.orders) == 0 {
		return fmt.Errorf("no orders")
	}

	lastOrder := &cafe.orders[len(cafe.orders)-1]

	if len(lastOrder.products) == 0 {
		return fmt.Errorf("no products inside order")
	}

	lastBeverage := lastOrder.products[len(lastOrder.products)-1]

	switch flavorType {
	case milk:
		lastOrder.products[len(lastOrder.products)-1] = flavors.NewMilk(lastBeverage)
	case mocha:
		lastOrder.products[len(lastOrder.products)-1] = flavors.NewMocha(lastBeverage)
	case soy:
		lastOrder.products[len(lastOrder.products)-1] = flavors.NewSoy(lastBeverage)
	case whip:
		lastOrder.products[len(lastOrder.products)-1] = flavors.NewWhip(lastBeverage)
	default:
		return fmt.Errorf("unknown flavor type")
	}
	return nil

}

func (cafe beverageCafe) OrderBeverage(beverageType int) error {
	if len(cafe.orders) == 0 {
		return fmt.Errorf("no orders")
	}

	lastOrder := &cafe.orders[len(cafe.orders)-1]

	var beverage beverages.IBeverage
	switch beverageType {
	case darkroast:
		beverage = beverages.NewDarkRoast(darkRoastCost)
	case houseblend:
		beverage = beverages.NewHouseBlend(houseBlendCost)
	case decaf:
		beverage = beverages.NewDecaf(decafCost)
	case espresso:
		beverage = beverages.NewEspresso(espressoCost)
	default:
		return fmt.Errorf("unknown beverage type")
	}

	lastOrder.products = append(lastOrder.products, beverage)

	return nil
}
