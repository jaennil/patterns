package flavors

import (
	"fmt"
)

func Print() {
	fmt.Println("---Flavors---")
	fmt.Printf("1. Milk %v$\n", milkCost)
	fmt.Printf("2. Mocha %v$\n", mochaCost)
	fmt.Printf("3. Soy %v$\n", soyCost)
	fmt.Printf("4. Whip %v$\n", whipCost)
}
