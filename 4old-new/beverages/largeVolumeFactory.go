package beverages

type LargeVolumeFactory struct {
}

// func NewLargeVolumeFactory() smallVolumeFactory {
// 	return smallVolumeFactory{}
// }

func (LargeVolumeFactory) CreateDarkRoast() darkRoast {
	return NewDarkRoast(.8)
}

func (LargeVolumeFactory) CreateHouseBlend() houseBlend {
	return NewHouseBlend(.8)
}

func (LargeVolumeFactory) CreateDecaf() decaf {
	return NewDecaf(.8)
}

func (LargeVolumeFactory) CreateEspresso() espresso {
	return NewEspresso(.8)
}
