package behaviors

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyOnWings struct{}

func (FlyOnWings) Fly() {
	fmt.Println("Я летаю на крыльях")
}

type NoFly struct{}

func (NoFly) Fly() {
	fmt.Println("Я не летаю")
}

type RemoteFly struct{}

func (RemoteFly) Fly() {
	fmt.Println("Я летаю на радиоуправлении")
}
