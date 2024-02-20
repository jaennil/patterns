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

const (
	small = iota + 1
	medium
	large
)

func printBeverageSizes() {
	fmt.Println("***Beverage Sizes***")
	fmt.Println("1. Small(0.4 L")
	fmt.Println("2. Medium(0.6 L")
	fmt.Println("3. Large(0.8 L")
	fmt.Println("********************")
	fmt.Println()
}

func printMenu() {
	fmt.Println("*****Menu*****")
	beverages.Print()
	flavors.Print()
	fmt.Println("**************")
	fmt.Println()
}

func printOrderCost(beverage beverages.IBeverage) {
	fmt.Println("---Order---")
	fmt.Println(beverage.Description())
	fmt.Println(beverage.Cost(), "$")
	fmt.Println("-----------")
	fmt.Println()
}

func printOperations(beverage beverages.IBeverage) {
	fmt.Println("*****Operations*****")
	fmt.Println("1. Print menu")
	fmt.Println("2. Order beverage")
	if beverage != nil {
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

func main() {

	printMenu()

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

		switch operationType {

		case printmenu:
			printMenu()
			fmt.Println()
			continue

		case createbeverage:
			beverages.Print()

			beverageType, err := readUserInput()
			if err != nil {
				fmt.Println("failed to scan your input. try again")
				continue
			}

			printBeverageSizes()

			beverageSize, err := readUserInput()
			if err != nil {
				fmt.Println("failed to scan your input. try again")
				continue
			}

			var factory beverages.IFactory

			switch beverageSize {
			case small:
				factory = beverages.SmallVolumeFactory{}
			case medium:
				factory = beverages.MediumVolumeFactory{}
			case large:
				factory = beverages.LargeVolumeFactory{}
			}

			switch beverageType {
			case darkroast:
				beverage = factory.CreateDarkRoast()
			case houseblend:
				beverage = factory.CreateHouseBlend()
			case decaf:
				beverage = factory.CreateDecaf()
			case espresso:
				beverage = factory.CreateEspresso()
			default:
				fmt.Println("unknown beverage type")
				continue
			}

		case addflavor:
			flavors.Print()

			flavorType, err := readUserInput()
			if err != nil {
				fmt.Println("failed to scan your input. try again")
				continue
			}

			switch flavorType {

			case milk:
				beverage = flavors.NewMilk(beverage)
			case mocha:
				beverage = flavors.NewMocha(beverage)
			case decaf:
				beverage = flavors.NewSoy(beverage)
			case espresso:
				beverage = flavors.NewWhip(beverage)
			}
		default:
			fmt.Println("unknown flavor type")
			continue
		}

		if beverage != nil {
			printOrderCost(beverage)
		}
	}
}
