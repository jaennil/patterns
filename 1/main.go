package main

import (
	"3/behaviors"
	"3/ducks"
	"fmt"
)

func main() {
	fmt.Println("Симуляция уток:")

	saxonyDuck := ducks.NewSaxonyDuck()
	saxonyDuck.Display()
	saxonyDuck.Quack()
	saxonyDuck.Fly()

	fmt.Println()

	decoyDuck := ducks.NewDecoyDuck()
	decoyDuck.Display()
	decoyDuck.Quack()
	decoyDuck.Fly()

	fmt.Println()

	rubberDuck := ducks.NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.Quack()
	rubberDuck.Fly()
	rubberDuck.SetFlyBehavior(behaviors.RemoteFly{})
	rubberDuck.Fly()

	fmt.Println()

	redheadDuck := ducks.NewRedheadDuck()
	redheadDuck.Display()
	redheadDuck.Quack()
	redheadDuck.SetQuackBehavior(behaviors.QuackLong{})
	redheadDuck.Quack()
	redheadDuck.Fly()
}
