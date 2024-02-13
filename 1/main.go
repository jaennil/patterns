package main

import (
	"1/behaviors"
	"1/ducks"
	"fmt"
)

func main() {

	fmt.Println("Симуляция уток:")

	saxonyDuck := ducks.SaxonyDuck{
		Duck: ducks.Duck{
			FlyBehavior:   &behaviors.FlyOnWings{},
			QuackBehavior: &behaviors.QuackRarely{},
		},
	}
	saxonyDuck.Display()
	saxonyDuck.PerformQuack()
	saxonyDuck.PerformFly()

	fmt.Println()

	rubberDuck := ducks.RubberDuck{
		Duck: ducks.Duck{
			FlyBehavior:   &behaviors.FlyNoFly{},
			QuackBehavior: &behaviors.QuackNoQuack{},
		},
	}
	rubberDuck.Display()
	rubberDuck.PerformQuack()
	rubberDuck.PerformFly()

	fmt.Println()

	decoyDuck := ducks.DecoyDuck{
		Duck: ducks.Duck{
			FlyBehavior:   &behaviors.FlyNoFly{},
			QuackBehavior: &behaviors.QuackLoud{},
		},
	}
	decoyDuck.Display()
	decoyDuck.PerformQuack()
	decoyDuck.PerformFly()
	decoyDuck.SetFlyBehavior(&behaviors.FlyRadio{})
	decoyDuck.PerformFly()

	fmt.Println()

	redheadDuck := ducks.RedheadDuck{
		Duck: ducks.Duck{
			FlyBehavior:   &behaviors.FlyOnWings{},
			QuackBehavior: &behaviors.QuackNoQuack{},
		},
	}
	redheadDuck.Display()
	redheadDuck.PerformQuack()
	redheadDuck.SetQuackBehavior(&behaviors.QuackLong{})
	redheadDuck.PerformQuack()
	redheadDuck.PerformFly()

}
