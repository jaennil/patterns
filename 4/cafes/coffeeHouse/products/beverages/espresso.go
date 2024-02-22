package beverages

type espresso struct {
	cost        float64
	description string
	size        float64
}

func NewEspresso(cost, size float64) espresso {
	return espresso{
		description: "Espresso",
		cost:        cost,
		size:        size,
	}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return espresso.description
}
