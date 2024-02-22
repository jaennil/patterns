package beverages

var darkRoastCost float64 = 1

type darkRoast struct {
	beverage
}

func NewDarkRoast(size float64) darkRoast {
	return darkRoast{
		beverage{
			cost:        darkRoastCost,
			description: "Dark Roast",
			size:        size,
		},
	}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.beverage.cost
}

func (darkRoast darkRoast) Description() string {
	return darkRoast.beverage.description
}
