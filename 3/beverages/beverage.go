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
	Description() string
}

type darkRoast struct {
	cost float64
	description string
}

func NewDarkRoast() darkRoast {
	return darkRoast{darkRoastCost, "Dark Roast"}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.cost
}

func (darkRoast darkRoast) Description() string {
	return darkRoast.description
}

type houseBlend struct {
	cost float64
	description string
}

func NewHouseBlend() houseBlend {
	return houseBlend{houseBlendCost, "House Blend"}
}

func (houseBlend houseBlend) Cost() float64 {
	return houseBlend.cost
}

func (houseBlend houseBlend) Description() string {
	return houseBlend.description
}

type espresso struct {
	cost float64
	description string
}

func NewEspresso() espresso {
	return espresso{espressoCost, "Espresso"}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return espresso.description
}

type decaf struct {
	cost float64
	description string
}

func NewDecaf() decaf {
	return decaf{decafCost, "Decaf"}
}

func (decaf decaf) Cost() float64 {
	return decaf.cost
}

func (decaf decaf) Description() string {
	return decaf.description
}
