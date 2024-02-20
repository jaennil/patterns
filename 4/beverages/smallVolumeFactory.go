package beverages

type SmallVolumeFactory struct {
}

// func NewSmallVolumeFactory() smallVolumeFactory {
// 	return smallVolumeFactory{}
// }

func (SmallVolumeFactory) CreateDarkRoast() darkRoast {
	return NewDarkRoast(.4)
}

func (SmallVolumeFactory) CreateHouseBlend() houseBlend {
	return NewHouseBlend(.4)
}

func (SmallVolumeFactory) CreateDecaf() decaf {
	return NewDecaf(.4)
}

func (SmallVolumeFactory) CreateEspresso() espresso {
	return NewEspresso(.4)
}
