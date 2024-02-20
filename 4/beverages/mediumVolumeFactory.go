package beverages

type MediumVolumeFactory struct {
}

// func NewSmallVolumeFactory() smallVolumeFactory {
// 	return smallVolumeFactory{}
// }

func (MediumVolumeFactory) CreateDarkRoast() darkRoast {
	return NewDarkRoast(.6)
}

func (MediumVolumeFactory) CreateHouseBlend() houseBlend {
	return NewHouseBlend(.6)
}

func (MediumVolumeFactory) CreateDecaf() decaf {
	return NewDecaf(.6)
}

func (MediumVolumeFactory) CreateEspresso() espresso {
	return NewEspresso(.6)
}
