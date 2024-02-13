package behaviors

import "fmt"

type QuackBehavior interface {
	Quack()
}

type QuackLoud struct {
}

func (quackLoud *QuackLoud) Quack() {
	fmt.Println("Я крякаю громко")
}

type QuackLong struct {
}

func (quackLong *QuackLong) Quack() {
	fmt.Println("Я крякаю протяжно")
}

type QuackNoQuack struct {
}

func (quackNoQuack *QuackNoQuack) Quack() {
	fmt.Println("Я не крякаю")
}

type QuackRarely struct {
}

func (quackRarely *QuackRarely) Quack() {
	fmt.Println("Я крякаю редко")
}