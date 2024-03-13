package main

type gumballMachine struct {
	gumballsAmount     int
	noCoinState        *noCoinState
	hasCoinState       *hasCoinState
	activatedState     *activatedState
	prizesAwardedState *prizesAwardedState
	state              state
}

func NewGumballMachine(gumballsAmount int) *gumballMachine {
	gumballMachine := &gumballMachine{gumballsAmount: gumballsAmount}

	gumballMachine.noCoinState = NewNoCoinState(gumballMachine)
	gumballMachine.hasCoinState = NewHasCoinState(gumballMachine)
	gumballMachine.activatedState = NewActivatedState(gumballMachine)
	gumballMachine.prizesAwardedState = NewPrizesAwardedState(gumballMachine)

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
	m.gumballsAmount -= 1
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

func (m *gumballMachine) GumballsAmount() int {
	return m.gumballsAmount;
}
