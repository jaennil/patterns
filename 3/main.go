package main

import (
	"3/cafes/coffeeHouse"
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
	fmt.Println()
	return input, err
}

func main() {
	cafe := coffeeHouse.NewCoffeeHouse()

	for {
		fmt.Print(cafe.Operations())

		operationType, err := readUserInput()
		if err != nil {
			fmt.Println("failed to scan your input. try again")
			continue
		}

		switch operationType {
		case printmenu:
			fmt.Println(cafe.Menu())
			continue
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

		order, err := cafe.Order()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(order)
	}
}
