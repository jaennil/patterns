package beverages

import "fmt"

func Print() {
	fmt.Println("---Beverages---")
	fmt.Printf("1. Dark Roast %v$\n", darkRoastCost)
	fmt.Printf("2. House Blend %v$\n", houseBlendCost)
	fmt.Printf("3. Decaf %v$\n", decafCost)
	fmt.Printf("4. Espresso %v$\n", espressoCost)
}

type IBeverage interface {
	Cost() float64
	Description() string
}

type beverage struct {
	cost float64
	description string
	size float64
}
