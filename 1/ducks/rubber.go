package ducks

import "fmt"

type RubberDuck struct {
	Duck
}

func (rubber *RubberDuck) Display() {
	fmt.Println("Резиновая утка")
}
