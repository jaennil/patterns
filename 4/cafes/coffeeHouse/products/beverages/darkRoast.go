package beverages

import "fmt"

type darkRoast struct {
	beverage
}

func NewDarkRoast(cost, size float64) darkRoast {
	return darkRoast{
		beverage{
			description: "Dark Roast",
			cost:        cost,
			size:        size,
		},
	}
}

func (darkRoast darkRoast) Cost() float64 {
	return darkRoast.cost
}

func (darkRoast darkRoast) Description() string {
	return fmt.Sprintf("%v %gL", darkRoast.description, darkRoast.size)
}
