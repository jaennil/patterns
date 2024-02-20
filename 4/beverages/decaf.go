package beverages

var decafCost float64 = 4

type decaf struct {
	beverage
}

func NewDecaf(size float64) decaf {
	return decaf{
		beverage{
			cost:        decafCost,
			description: "Decaf",
			size:        size,
		},
	}
}

func (decaf decaf) Cost() float64 {
	return decaf.cost
}

func (decaf decaf) Description() string {
	return decaf.description
}
