package beverages

type darkRoast struct {
	description string
	cost        float64
	size        float64
}

func NewDarkRoast(cost, size float64) darkRoast {
	return darkRoast{
		description: "DarkRoast",
		cost:        cost,
		size:        size,
	}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.cost
}

func (darkRoast darkRoast) Description() string {
	return darkRoast.description
}
