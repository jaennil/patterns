package beverages

type DarkRoast struct {
	IBeverage
}

func (darkRoast DarkRoast) Cost() int {
	return 1
}

func (darkRoast DarkRoast) GetDescription() string {
	return "Dark Roast"
}
