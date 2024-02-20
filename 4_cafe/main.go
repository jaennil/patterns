package main

import (
	"3/cafes"
	"fmt"
)

type operationType int

const (
	printmenu operationType = iota + 1
	orderbeverage
	addflavor
)

func readUserInput() (int, error) {
	fmt.Print(">> ")
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}

func printOperations() {
	fmt.Println("*****Operations*****")
	fmt.Println("1. Print menu")
	fmt.Println("2. Order beverage")
	// if beverage != nil {
	// 	fmt.Println("3. Add flavor")
	// }
	fmt.Println("any. Exit")
	fmt.Println("********************")
}

func main() {

	cafe := cafes.NewCafe()

	for {
		printOperations()

		input, err := readUserInput()
		if err != nil {
			fmt.Println("failed to scan your input. try again")
			continue
		}

		switch operationType(input) {
		case printmenu:
			cafes.PrintMenu()
			fmt.Println()
		case orderbeverage:
			beverageType, err := readUserInput()
			if err != nil {
				fmt.Println(err)
			}

			return cafe.OrderBeverage(beverageType)
		case addflavor:
			flavors.Print()

			flavorType, err := readUserInput()
			if err != nil {
				return nil, err
			}

			return addFlavor(beverage, flavorType)
		}

		cafe.PrintOrderCost()
	}
}
