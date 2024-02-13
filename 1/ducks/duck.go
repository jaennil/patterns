package ducks

import (
	"1/behaviors"
	"fmt"
)

type Duck struct {
	FlyBehavior   behaviors.FlyBehavior
	QuackBehavior behaviors.QuackBehavior
}

func (duck *Duck) Display() {
	fmt.Println("Я утка")
}

func (duck *Duck) Quack() {
	fmt.Println("Я крякаю")
}

func (duck *Duck) Swim() {
	fmt.Println("Я плаваю")
}

func (duck *Duck) PerformFly() {
	duck.FlyBehavior.Fly()
}

func (duck *Duck) PerformQuack() {
	duck.QuackBehavior.Quack()
}

func (duck *Duck) SetFlyBehavior(flyBehavior behaviors.FlyBehavior) {
	duck.FlyBehavior = flyBehavior
}

func (duck *Duck) SetQuackBehavior(quackBehavior behaviors.QuackBehavior) {
	duck.QuackBehavior = quackBehavior
}
