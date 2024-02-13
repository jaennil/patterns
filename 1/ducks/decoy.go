package ducks

import "fmt"

type DecoyDuck struct {
	Duck
}

func (decoy *DecoyDuck) Display() {
	fmt.Println("Утка-приманка")
}

