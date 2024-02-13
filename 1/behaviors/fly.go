package behaviors

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyOnWings struct {
}

func (flyOnWings *FlyOnWings) Fly() {
	fmt.Println("Я летаю на крыльях")
}

type FlyNoFly struct {
}

func (flyNoFly *FlyNoFly) Fly() {
	fmt.Println("Я не летаю")
}

type FlyRadio struct {
}

func (flyRadio *FlyRadio) Fly() {
	fmt.Println("Я летаю на радиоуправлении")
}