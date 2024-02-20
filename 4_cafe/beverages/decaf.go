package beverages

var decafCost float64 = 4

type decaf struct {
	cost        float64
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
