package beverages

import "fmt"

var (
	darkRoastCost float64 = 1
	houseBlendCost float64 = 2
	espressoCost float64 = 3
	decafCost float64 = 4
)

func Print() {
	fmt.Println("---Beverages---")
	fmt.Printf("1. Dark Roast %v$\n", darkRoastCost)
	fmt.Printf("2. House Blend %v$\n", houseBlendCost)
	fmt.Printf("3. Decaf %v$\n", decafCost)
	fmt.Printf("4. Espresso %v$\n", espressoCost)
}

type IBeverage interface {
	Cost() float64
}

type darkRoast struct {
	cost float64
}

func NewDarkRoast() darkRoast {
	return darkRoast{darkRoastCost}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.cost
}

type houseBlend struct {
	cost float64
}

func NewHouseBlend() houseBlend {
	return houseBlend{houseBlendCost}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

type espresso struct {
	cost float64
}

func NewEspresso() espresso {
	return espresso{espressoCost}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

type decaf struct {
	cost float64
}

func NewDecaf() decaf {
	return decaf{decafCost}
}

func (decaf decaf) Cost() float64 {
	return decaf.cost
}
