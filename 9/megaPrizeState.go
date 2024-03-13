package main

import (
	"fmt"
)

type megaPrizeState struct {
	gumballMachine *gumballMachine
}

func NewMegaPrizeState(gumballMachine *gumballMachine) *megaPrizeState {
	return &megaPrizeState{gumballMachine}
}

func (megaPrizeState) insertCoin() {
	fmt.Println("cant insert coin. pull the prize first")
}

func (s *megaPrizeState) pullCoin() {
	fmt.Println("cant pull coin. there is no coin inserted")
}

func (s *megaPrizeState) activate() {
	fmt.Println("cant activate the machine. it is already activated")
}

func (s *megaPrizeState) pullPrize() {
	fmt.Println("you lucky! you got mega prize: toy and double gumball")
	s.gumballMachine.decreaseToy()
	s.gumballMachine.decreaseGumball()
	s.gumballMachine.decreaseGumball()
	if s.gumballMachine.GumballsAmount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.NoCoinState())
		return
	}

	s.gumballMachine.setState(s.gumballMachine.PrizesAwardedState())
}
