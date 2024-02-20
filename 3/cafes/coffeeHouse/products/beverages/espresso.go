package beverages

type espresso struct {
	cost        float64
	description string
}

func NewEspresso(cost float64) espresso {
	return espresso{
		description: "Espresso",
		cost:        cost,
	}
}

func (espresso espresso) Cost() float64 {
	return espresso.cost
}

func (espresso espresso) Description() string {
	return espresso.description
}
