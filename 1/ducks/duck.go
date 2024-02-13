package ducks

import "3/behaviors"

type Duck struct {
	flyBehavior behaviors.FlyBehavior
	quackBehavior behaviors.QuackBehavior
}

func (duck *Duck) SetFlyBehavior(flyBehavior behaviors.FlyBehavior) {
	duck.flyBehavior = flyBehavior
}

func (duck Duck) Fly() {
	duck.flyBehavior.Fly()
}

func (duck *Duck) SetQuackBehavior(quackBehavior behaviors.QuackBehavior) {
	duck.quackBehavior = quackBehavior
}

func (duck Duck) Quack() {
	duck.quackBehavior.Quack()
}
