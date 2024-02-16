package beverages

var darkRoastCost float64 = 1

type darkRoast struct {
	cost        float64
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
