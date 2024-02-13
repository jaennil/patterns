package ducks

import (
	"3/behaviors"
	"fmt"
)

type DecoyDuck struct {
	Duck
}

func NewDecoyDuck() DecoyDuck {
	return DecoyDuck{Duck{behaviors.NoFly{}, behaviors.NoQuack{}}}
}

func (DecoyDuck) Display() {
	fmt.Println("Резиновая утка")
}
