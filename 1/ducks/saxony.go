package ducks

import (
	"3/behaviors"
	"fmt"
)

type SaxonyDuck struct {
	Duck
}

func NewSaxonyDuck() SaxonyDuck {
	return SaxonyDuck{Duck{behaviors.FlyOnWings{}, behaviors.QuackRare{}}}
}

func (SaxonyDuck) Display() {
	fmt.Println("Саксонская утка")
}
