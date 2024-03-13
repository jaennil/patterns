package main

type state interface {
	insertCoin()
	pullCoin()
	activate()
	pullPrize()
}
