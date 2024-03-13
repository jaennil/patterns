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

	fmt.Print("toys amount: ")
	var toysAmount int
	_, err = fmt.Scanln(&toysAmount)
	if err != nil {
		fmt.Println("wrong input")
		return
	}

	fmt.Println()

	gumballMachine := NewGumballMachine(gumballsAmount, toysAmount)

	var action int

	for {
		fmt.Println(`
1. insert coin
2. pull coin
3. activate machine
4. pull prize
5. print left prizes
6. ask for refill`)
		_, err = fmt.Scanln(&action)
		if err != nil {
			fmt.Println("wrong input")
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
		case 5:
			gumballsLeft, toysLeft, err := gumballMachine.leftPrizes()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("gumballs left: %v\ntoysLeft: %v\n", gumballsLeft, toysLeft)
		case 6:
			err = gumballMachine.askForRefill()
			if err == nil {
				fmt.Println("asked to refill")
				continue
			}

			fmt.Println("faled to ask for refill: ", err)
		}
	}
}
