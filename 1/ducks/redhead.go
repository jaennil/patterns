package ducks

import (
	"3/behaviors"
	"fmt"
)

type RedheadDuck struct {
	Duck
}

func NewRedheadDuck() RedheadDuck {
	return RedheadDuck{Duck{behaviors.FlyOnWings{}, behaviors.NoQuack{}}}
}

func (RedheadDuck) Display() {
	fmt.Println("Красноголовая утка")
}
