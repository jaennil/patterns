package main

import "fmt"

type noCoinState struct {
	gumballMachine *gumballMachine
}

func NewNoCoinState(gumballMachine *gumballMachine) *noCoinState {
	return &noCoinState{gumballMachine}
}

func (s *noCoinState) insertCoin() {
	fmt.Println("coin inserted")
	s.gumballMachine.setState(s.gumballMachine.HasCoinState())

}
func (noCoinState) pullCoin() {
	fmt.Println("cant pull coin. you dont have coin inserted")
}
func (noCoinState) activate() {
	fmt.Println("cant activate the machine. you need to insert a coin first")
}
func (noCoinState) pullPrize() {
	fmt.Println("cant pull the prize. you need to pay first");
}
