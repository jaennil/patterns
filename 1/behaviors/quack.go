package behaviors

import "fmt"

type QuackBehavior interface {
	Quack()
}

type QuackRare struct{}

func (QuackRare) Quack() {
	fmt.Println("Я крякаю редко")
}

type NoQuack struct{}

func (NoQuack) Quack() {
	fmt.Println("Я не крякаю")
}

type QuackLoud struct{}

func (QuackLoud) Quack() {
	fmt.Println("Я крякаю громко")
}

type QuackLong struct{}

func (QuackLong) Quack() {
	fmt.Println("Я крякаю протяжно")
}
