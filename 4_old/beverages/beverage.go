package beverages

type IBeverage interface {
	GetDescription() string
	Cost() int
}
