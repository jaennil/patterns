package main

import (
	"fmt"
)

type doublePrizeState struct {
	gumballMachine *gumballMachine
}

func NewDoublePrizeState(gumballMachine *gumballMachine) *doublePrizeState {
	return &doublePrizeState{gumballMachine}
}

func (doublePrizeState) insertCoin() {
	fmt.Println("cant insert coin. pull the prize first")
}

func (s *doublePrizeState) pullCoin() {
	fmt.Println("cant pull coin. there is no coin inserted")
}

func (s *doublePrizeState) activate() {
	fmt.Println("cant activate the machine. it is already activated")
}

func (s *doublePrizeState) pullPrize() {
	fmt.Println("you got gumball")
	s.gumballMachine.decreaseGumball()
	if s.gumballMachine.GumballsAmount() == 0 {
		s.gumballMachine.setState(s.gumballMachine.PrizesAwardedState())
		return
	}

	fmt.Println("you lucky! you got second gumball for free")
	s.gumballMachine.decreaseGumball()

	if s.gumballMachine.GumballsAmount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.NoCoinState())
		return
	}

	fmt.Println("all prizes awarded")
	s.gumballMachine.setState(s.gumballMachine.PrizesAwardedState())
}
