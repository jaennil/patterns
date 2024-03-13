package main

import "fmt"

func main() {
	fmt.Print("gumballs amount: ")
	var gumballsAmount int
	_, err := fmt.Scanln(&gumballsAmount)
	if err != nil {
		fmt.Println("wrong input")
		return
	}

	fmt.Println()

	gumballMachine := NewGumballMachine(gumballsAmount)

	var action int

	for {
		fmt.Println(`1. insert coin
2. pull coin
3. activate machine
4. pull prize`)
		_, err = fmt.Scanln(&action)
		if err != nil {
			fmt.Println("wring input")
			return
		}

		switch action {
		case 1:
			gumballMachine.insertCoin()
		case 2:
			gumballMachine.pullCoin()
		case 3:
			gumballMachine.activate()
		case 4:
			gumballMachine.pullPrize()
		}
	}
}
