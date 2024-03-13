package main

import "fmt"

type gumballMachine struct {
	gumballsAmount     int
	toysAmount         int
	noCoinState        *noCoinState
	hasCoinState       *hasCoinState
	activatedState     *activatedState
	prizesAwardedState *prizesAwardedState
	doublePrizeState   *doublePrizeState
	toyPrizeState      *toyPrizeState
	megaPrizeState     *megaPrizeState
	state              state
}

func NewGumballMachine(gumballsAmount, toysAmount int) *gumballMachine {
	gumballMachine := &gumballMachine{gumballsAmount: gumballsAmount, toysAmount: toysAmount}

	gumballMachine.noCoinState = NewNoCoinState(gumballMachine)
	gumballMachine.hasCoinState = NewHasCoinState(gumballMachine)
	gumballMachine.activatedState = NewActivatedState(gumballMachine)
	gumballMachine.prizesAwardedState = NewPrizesAwardedState(gumballMachine)
	gumballMachine.doublePrizeState = NewDoublePrizeState(gumballMachine)
	gumballMachine.toyPrizeState = NewToyPrizeState(gumballMachine)
	gumballMachine.megaPrizeState = NewMegaPrizeState(gumballMachine)

	if gumballsAmount > 0 {
		gumballMachine.state = gumballMachine.noCoinState
	} else {
		gumballMachine.state = gumballMachine.prizesAwardedState
	}

	return gumballMachine
}

func (gumballMachine *gumballMachine) insertCoin() {
	gumballMachine.state.insertCoin()
}

func (gumballMachine *gumballMachine) pullCoin() {
	gumballMachine.state.pullCoin()
}

func (gumballMachine *gumballMachine) activate() {
	gumballMachine.state.activate()
}

func (gumballMachine *gumballMachine) pullPrize() {
	gumballMachine.state.pullPrize()
}

func (gumballMachine *gumballMachine) setState(state state) {
	gumballMachine.state = state
}

func (m *gumballMachine) decreaseGumball() {
	if m.gumballsAmount > 0 {
		m.gumballsAmount -= 1
	}
}

func (m *gumballMachine) decreaseToy() {
	if m.toysAmount > 0 {
		m.toysAmount -= 1
	}
}

func (gumballMachine *gumballMachine) HasCoinState() *hasCoinState {
	return gumballMachine.hasCoinState
}

func (gumballMachine *gumballMachine) NoCoinState() *noCoinState {
	return gumballMachine.noCoinState
}

func (gumballMachine *gumballMachine) ActivatedState() *activatedState {
	return gumballMachine.activatedState
}

func (gumballMachine *gumballMachine) PrizesAwardedState() *prizesAwardedState {
	return gumballMachine.prizesAwardedState
}

func (gumballMachine *gumballMachine) DoublePrizeState() *doublePrizeState {
	return gumballMachine.doublePrizeState
}

func (gumballMachine *gumballMachine) ToyPrizeState() *toyPrizeState {
	return gumballMachine.toyPrizeState
}

func (gumballMachine *gumballMachine) MegaPrizeState() *megaPrizeState {
	return gumballMachine.megaPrizeState
}

func (m *gumballMachine) GumballsAmount() int {
	return m.gumballsAmount
}

func (m *gumballMachine) ToysAmount() int {
	return m.toysAmount
}

func (m *gumballMachine) leftPrizes() (gumballsAmount, toysAmount int, err error) {
	if m.state == m.noCoinState {
		return m.gumballsAmount, m.toysAmount, nil
	}

	return -1, -1, fmt.Errorf("left prizes can be known only in noCoinState")
}
