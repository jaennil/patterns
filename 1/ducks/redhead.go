package ducks

import "fmt"

type RedheadDuck struct {
	Duck
}

func (redheadDuck *RedheadDuck) Display() {
	fmt.Println("Красноголовая утка")
}