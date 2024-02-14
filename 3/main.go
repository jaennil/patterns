package main

import (
	"3/beverages"
	"3/flavors"
	"fmt"
)

const (
	milk = iota + 1
	mocha
	soy
	whip
)

const (
	printmenu = iota + 1
	createbeverage
	addflavor
)

const (
	darkroast = iota + 1
	houseblend
	decaf
	espresso
)

func printMenu() {
	fmt.Println("*****Menu*****")
	beverages.Print()
	flavors.Print()
	fmt.Println("**************")
}

func printOrderCost(beverage beverages.IBeverage) {
	fmt.Println("---Total Order Cost---")
	fmt.Println(beverage.Cost(), "$")
	fmt.Println("----------------------")
	fmt.Println()
}

func printOperations(beverage beverages.IBeverage) {
	fmt.Println("*****Operations*****")
	fmt.Println("1. Print menu")
	fmt.Println("2. Order beverage")
	if beverage != nil{
		fmt.Println("3. Add flavor")
	}
	fmt.Println("any. Exit")
	fmt.Println("********************")
}

func readUserInput() (int, error) {
	fmt.Print(">> ")
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}

func createBeverage(beverageType int) (beverages.IBeverage, error) {
	var beverage beverages.IBeverage

	switch beverageType {
	case darkroast:
		beverage = beverages.NewDarkRoast()
	case houseblend:
		beverage = beverages.NewHouseBlend()
	case decaf:
		beverage = beverages.NewDecaf()
	case espresso:
		beverage = beverages.NewEspresso()
	default:
		return nil, fmt.Errorf("unknown beverage type")
	}

	return beverage, nil
}

func handleOperation(operationType int, beverage beverages.IBeverage) (beverages.IBeverage, error) {
	switch operationType {

	case printmenu:
		printMenu()
		fmt.Println()
		return beverage, nil

	case createbeverage:
		beverages.Print()

		beverageType, err := readUserInput()
		if err != nil {
			return nil, err
		}

		return createBeverage(beverageType)

	case addflavor:
		flavors.Print()

		flavorType, err := readUserInput()
		if err != nil {
			return nil, err
		}

		return addFlavor(beverage, flavorType)
	}

	return nil, fmt.Errorf("unknown operation")
}

func addFlavor(beverage beverages.IBeverage, flavorType int) (beverages.IBeverage, error) {

	switch flavorType {

	case milk:
		return flavors.NewMilk(beverage), nil
	case mocha:
		return flavors.NewMocha(beverage), nil
	case decaf:
		return flavors.NewSoy(beverage), nil
	case espresso:
		return flavors.NewWhip(beverage), nil
	}

	return nil, fmt.Errorf("unknown flavor type")
}

func main() {

	printMenu()
	fmt.Println()

	var beverage beverages.IBeverage
	var operationType int
	var err error

	for {
		printOperations(beverage)

		operationType, err = readUserInput()
		if err != nil {
			fmt.Println("failed to scan your input. try again")
			continue
		}

		beverage, err = handleOperation(operationType, beverage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if beverage != nil {
			printOrderCost(beverage)
		}
	}
}
