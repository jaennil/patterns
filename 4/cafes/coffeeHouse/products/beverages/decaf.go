package beverages

type decaf struct {
	description string
	cost        float64
	size        float64
}

func NewDecaf(cost, size float64) decaf {
	return decaf{
		description: "Decaf",
		cost:        cost,
		size:        size,
	}
}

func (decaf decaf) Cost() float64 {
	return decaf.cost
}

func (decaf decaf) Description() string {
	return decaf.description
}
