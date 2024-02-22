package beverages

type IFactory interface {
	// CreateDarkRoast(cost float64) darkRoast
	// CreateHouseBlend(cost float64) houseBlend
	// CreateDecaf(cost float64) decaf
	CreateBeverage(beverageType int, cost float64) (IBeverage, error)
}
