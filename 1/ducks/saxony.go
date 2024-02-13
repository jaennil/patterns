package ducks

import "fmt"

type SaxonyDuck struct {
	Duck
}

func (saxonyDuck *SaxonyDuck) Display() {
	fmt.Println("Саксонская утка")
}
