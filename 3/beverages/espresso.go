package beverages

var espressoCost float64 = 3

type espresso struct {
	cost        float64
	description string
}

func NewEspresso() espresso {
	return espresso{espressoCost, "Espresso"}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return espresso.description
}
