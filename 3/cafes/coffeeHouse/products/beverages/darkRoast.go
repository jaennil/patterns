package beverages

type darkRoast struct {
	description string
	cost        float64
}

func NewDarkRoast(cost float64) darkRoast {
	return darkRoast{
		description: "DarkRoast",
		cost:        cost,
	}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.cost
}

func (darkRoast darkRoast) Description() string {
	return darkRoast.description
}
