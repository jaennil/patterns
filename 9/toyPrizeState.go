package main

import (
	"fmt"
)

type toyPrizeState struct {
	gumballMachine *gumballMachine
}

func NewToyPrizeState(gumballMachine *gumballMachine) *toyPrizeState {
	return &toyPrizeState{gumballMachine}
}

func (toyPrizeState) insertCoin() {
	fmt.Println("cant insert coin. pull the prize first")
}

func (s *toyPrizeState) pullCoin() {
	fmt.Println("cant pull coin. there is no coin inserted")
}

func (s *toyPrizeState) activate() {
	fmt.Println("cant activate the machine. it is already activated")
}

func (s *toyPrizeState) pullPrize() {
	fmt.Println("you lucky! you got toy prize")
	s.gumballMachine.decreaseToy()
	s.gumballMachine.setState(s.gumballMachine.NoCoinState())
}
