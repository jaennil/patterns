package beverages

type IBeverage interface {
	Cost() float64
	Description() string
}

type beverage struct {
	description string
	cost        float64
	size        float64
}
