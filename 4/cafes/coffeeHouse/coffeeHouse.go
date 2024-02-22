package coffeeHouse

import (
	"3/cafes"
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

var beverageCost = map[int]map[int]float64{
	beverages.Darkroast: {
		small:  darkRoastCost * small,
		medium: darkRoastCost * medium,
		large:  darkRoastCost * large,
	},
	beverages.Houseblend: {
		small:  houseBlendCost * small,
		medium: houseBlendCost * medium,
		large:  houseBlendCost * large,
	},
	beverages.Decaf: {
		small:  decafCost * small,
		medium: decafCost * medium,
		large:  decafCost * large,
	},
	beverages.Espresso: {
		small:  espressoCost * small,
		medium: espressoCost * medium,
		large:  espressoCost * large,
	},
}

const (
	milk = iota + 1
	mocha
	soy
	whip
)

const (
	small = iota + 1
	medium
	large
)

var operations = `
*****Operations*****
1. Print menu
2. Order beverage
3. Add flavor
4. New Order
any. Exit
********************
`

var beveragesOperations = `---Beverages---
1. Dark Roast
2. House Blend
3. Decaf
4. Espresso
`

var beverageSizes = `---Beverage Sizes---
1. Small (0.4 L)
2. Medium (0.6 L)
3. Large (0.8 L)
`

var beveragesMenu = fmt.Sprintf(
	`---Beverages---
1. Dark Roast Small(0.4 L) %v $
2. Dark Roast Medium(0.6 L) %v $
3. Dark Roast Large(0.8 L) %v $
4. House Blend Small(0.4 L) %v $
5. House Blend Medium(0.6 L) %v $
6. House Blend Large(0.8 L) %v $
7. Decaf Small(0.4 L) %v $
8. Decaf Medium(0.6 L) %v $
9. Decaf Large(0.8 L) %v $
10. Espresso Small(0.4 L) %v $
11. Espresso Medium(0.6 L) %v $
12. Espresso Large(0.8 L) %v $`,
	beverageCost[beverages.Darkroast][small],
	beverageCost[beverages.Darkroast][medium],
	beverageCost[beverages.Darkroast][large],
	beverageCost[beverages.Houseblend][small],
	beverageCost[beverages.Houseblend][medium],
	beverageCost[beverages.Houseblend][large],
	beverageCost[beverages.Decaf][small],
	beverageCost[beverages.Decaf][medium],
	beverageCost[beverages.Decaf][large],
	beverageCost[beverages.Espresso][small],
	beverageCost[beverages.Espresso][medium],
	beverageCost[beverages.Espresso][large])

var flavorsMenu = fmt.Sprintf(
	`---Flavors---
1. Milk %v$
2. Mocha %v$
3. Soy %v$
4. Whip %v$
`,
	milkCost, mochaCost, soyCost, whipCost)

var menu = fmt.Sprintf(`
*****Menu*****%v%v**************`, beveragesMenu, flavorsMenu)

type coffeeHouse struct {
	cafes.Cafe
}

func NewCoffeeHouse() coffeeHouse {
	return coffeeHouse{}
}

func (coffeeHouse) Operations() string {
	return operations
}

func (coffeeHouse) BeveragesOperations() string {
	return beveragesOperations
}

func (coffeeHouse) BeverageSizes() string {
	return beverageSizes
}

func (coffeeHouse) BeveragesMenu() string {
	return beveragesMenu
}

func (coffeeHouse) FlavorsMenu() string {
	return flavorsMenu
}

func (coffeeHouse) Menu() string {
	return menu
}

func (cafe coffeeHouse) AddFlavor(flavorType int) error {

	if len(cafe.Orders) == 0 {
		return fmt.Errorf("no orders")
	}

	lastOrder := &cafe.Orders[len(cafe.Orders)-1]

	if len(lastOrder.Products) == 0 {
		return fmt.Errorf("no products inside order")
	}

	lastBeverage := &lastOrder.Products[len(lastOrder.Products)-1]

	switch flavorType {
	case milk:
		*lastBeverage = flavors.NewMilk(*lastBeverage)
	case mocha:
		*lastBeverage = flavors.NewMocha(*lastBeverage)
	case soy:
		*lastBeverage = flavors.NewSoy(*lastBeverage)
	case whip:
		*lastBeverage = flavors.NewWhip(*lastBeverage)
	default:
		return fmt.Errorf("unknown flavor type")
	}
	return nil

}

func (cafe coffeeHouse) OrderBeverage(beverageType int, size int) error {
	if len(cafe.Orders) == 0 {
		return fmt.Errorf("no orders")
	}

	lastOrder := &cafe.Orders[len(cafe.Orders)-1]

	var factory beverages.IFactory

	switch size {
	case small:
		factory = beverages.SmallVolumeFactory{}
	case medium:
		factory = beverages.MediumVolumeFactory{}
	case large:
		factory = beverages.LargeVolumeFactory{}
	}

	beverage, err := factory.CreateBeverage(beverageType, beverageCost[beverageType][size])
	if err != nil {
		return err
	}

	lastOrder.Products = append(lastOrder.Products, beverage)

	return nil
}
