package main

import (
	"3/cafes"
	"fmt"
)

const (
	printmenu = iota + 1
	createbeverage
	addflavor
	neworder
)

func readUserInput() (int, error) {
	fmt.Print(">> ")
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}

func main() {
	cafe := cafes.NewBeverageCafe()

	for {
		cafe.PrintOperations()

		operationType, err := readUserInput()
		if err != nil {
			fmt.Println("failed to scan your input. try again")
			continue
		}

		switch operationType {
		case printmenu:
			cafe.PrintMenu()
			fmt.Println()
		case createbeverage:
			fmt.Print(cafe.BeveragesMenu())

			beverageType, err := readUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = cafe.OrderBeverage(beverageType)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case addflavor:
			fmt.Print(cafe.FlavorsMenu())

			flavorType, err := readUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}

			cafe.AddFlavor(flavorType)
		case neworder:
			cafe.NewOrder()
			continue
		}

		err = cafe.PrintOrder()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
