package main

import (
	"fmt"
)

type hasCoinState struct {
	gumballMachine *gumballMachine
}

func NewHasCoinState(gumballMachine *gumballMachine) *hasCoinState {
	return &hasCoinState{gumballMachine}
}

func (hasCoinState) insertCoin() {
	fmt.Println("cant insert coin. you already inserted another coin")
}

func (s *hasCoinState) pullCoin() {
	fmt.Println("coin pulled")
	s.gumballMachine.setState(s.gumballMachine.NoCoinState())
}

func (s *hasCoinState) activate() {
	fmt.Println("activated")
	s.gumballMachine.setState(s.gumballMachine.ActivatedState())
}

func (hasCoinState) pullPrize() {
	fmt.Println("cant pull prize. you need to activate machine first")
}
