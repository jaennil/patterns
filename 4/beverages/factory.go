package beverages

type IFactory interface {
	CreateDarkRoast() darkRoast
	CreateHouseBlend() houseBlend
	CreateDecaf() decaf
	CreateEspresso() espresso
}
