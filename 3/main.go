package main

import (
	"3/beverages"
	"3/flavors"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("1 DarkRoast, price 1")
	fmt.Println("2 decaf, price 2")
	fmt.Println("3 espresso, price 3")
	fmt.Println("4 house blend, price 4")
	fmt.Println("other - exit")

	var beverage beverages.IBeverage

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	beverageInput, _ := strconv.Atoi(scanner.Text())
	switch beverageInput {
	case 1:
		beverage = &beverages.DarkRoast{}
		fmt.Println("totalCost:", beverage.Cost())
		fmt.Println("description:", beverage.GetDescription())
	case 2:
		beverage = &beverages.Decaf{}
		fmt.Println("totalCost:", beverage.Cost())
		fmt.Println("description:", beverage.GetDescription())
	case 3:
		beverage = &beverages.Espresso{}
		fmt.Println("totalCost:", beverage.Cost())
		fmt.Println("description:", beverage.GetDescription())
	case 4:
		beverage = &beverages.HouseBlend{}
		fmt.Println("totalCost:", beverage.Cost())
		fmt.Println("description:", beverage.GetDescription())
	default:
		return
	}

	fmt.Println("1 milk, price 1")
	fmt.Println("2 mocha, price 2")
	fmt.Println("3 soy, price 3")
	fmt.Println("4 whip, price 4")
	fmt.Println("other - exit")

loop:
	for {
		scanner.Scan()
		flavorInput, _ := strconv.Atoi(scanner.Text())
		switch flavorInput {
		case 1:
			beverage = &flavors.Milk{Beverage: beverage}
			fmt.Println("totalCost:", beverage.Cost())
			fmt.Println("description:", beverage.GetDescription())
		case 2:
			beverage = &flavors.Mocha{Beverage: beverage}
			fmt.Println("totalCost:", beverage.Cost())
			fmt.Println("description:", beverage.GetDescription())
		case 3:
			beverage = &flavors.Soy{Beverage: beverage}
			fmt.Println("totalCost:", beverage.Cost())
			fmt.Println("description:", beverage.GetDescription())
		case 4:
			beverage = &flavors.Whip{Beverage: beverage}
			fmt.Println("totalCost:", beverage.Cost())
			fmt.Println("description:", beverage.GetDescription())
		default:
			fmt.Println("exit")
			break loop
		}
	}
}
