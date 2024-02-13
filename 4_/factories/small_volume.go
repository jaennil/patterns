package factories

type SmallVolumeFactory struct{}

func (factory SmallVolumeFactory) CreateBase() Base {
	return Water
}