package ducks

import (
	"3/behaviors"
	"fmt"
)

type RubberDuck struct {
	Duck
}

func NewRubberDuck() RubberDuck {
	return RubberDuck{Duck{behaviors.NoFly{}, behaviors.QuackLoud{}}}
}

func (RubberDuck) Display() {
	fmt.Println("Утка-приманка")
}
