package beverages

type Decaf struct {
	IBeverage
}

func (decaf *Decaf) Cost() int {
	return 2
}

func (decaf *Decaf) GetDescription() string {
	return "Decaf"
}
