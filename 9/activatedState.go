package main

import (
	"fmt"
)

type activatedState struct {
	gumballMachine *gumballMachine
}

func NewActivatedState(gumballMachine *gumballMachine) *activatedState {
	return &activatedState{gumballMachine}
}

func (activatedState) insertCoin() {
	fmt.Println("cant insert coin. pull the prize first")
}

func (activatedState) pullCoin() {
	fmt.Println("cant pull coin. there is no coin inserted")
}

func (activatedState) activate() {
	fmt.Println("cant activate the machine. it is already activated")
}

func (s *activatedState) pullPrize() {
	s.gumballMachine.decreaseGumball()
	fmt.Println("you pulled the prize")
	if s.gumballMachine.GumballsAmount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.NoCoinState())
	} else {
		fmt.Println("all prizes awarded")
		s.gumballMachine.setState(s.gumballMachine.PrizesAwardedState())
	}
}
