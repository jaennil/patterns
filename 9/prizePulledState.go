package main

import "fmt"

type prizesAwardedState struct {
	gumballMachine *gumballMachine
}

func NewPrizesAwardedState(gumballMachine *gumballMachine) *prizesAwardedState {
	return &prizesAwardedState{gumballMachine}
}

func (prizesAwardedState) insertCoin() {
	fmt.Println("cant insert coin. all prizes are awarded")
}
func (prizesAwardedState) pullCoin() {
	fmt.Println("cant pull coin. no coin inserted")
}
func (prizesAwardedState) activate() {
	fmt.Println("cant activate machine. no coin inserted")
}
func (prizesAwardedState) pullPrize() {
	fmt.Println("cant pull prize. machine wasnt activated");
}
