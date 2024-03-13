package main

import (
	"fmt"
	"math/rand"
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

	toyRand := rand.Float32()
	doubleRand := rand.Float32()

	if toyRand < 0.15 && doubleRand < 0.1 && s.gumballMachine.GumballsAmount() > 1 && s.gumballMachine.ToysAmount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.MegaPrizeState())
		return
	}

	if toyRand < 0.15 && s.gumballMachine.ToysAmount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.ToyPrizeState())
		return
	}

	if doubleRand < 0.1 && s.gumballMachine.GumballsAmount() > 1 {
		s.gumballMachine.setState(s.gumballMachine.DoublePrizeState())
		return
	}

	s.gumballMachine.setState(s.gumballMachine.ActivatedState())
}

func (hasCoinState) pullPrize() {
	fmt.Println("cant pull prize. you need to activate machine first")
}
