package main

import (
	"3/cafes/coffeeHouse"
	"fmt"
)

const (
	printmenu = iota + 1
	orderbeverage
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
		fmt.Println(cafe.Operations())

		operationType, err := readUserInput()
		if err != nil {
			fmt.Println("failed to scan your input. try again")
			continue
		}

		switch operationType {
		case printmenu:
			fmt.Println(cafe.Menu())
			continue
		case orderbeverage:
			fmt.Println(cafe.BeveragesOperations())

			beverageType, err := readUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(cafe.BeverageSizes())

			beverageSize, err := readUserInput()
			if err != nil {
				fmt.Println("failed to scan your input. try again")
				continue
			}

			err = cafe.OrderBeverage(beverageType, beverageSize)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case addflavor:
			fmt.Println(cafe.FlavorsMenu())

			flavorType, err := readUserInput()
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = cafe.AddFlavor(flavorType)
			if err != nil {
				fmt.Println(err)
				continue
			}
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
