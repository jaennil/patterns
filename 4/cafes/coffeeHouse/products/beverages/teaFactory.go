package beverages

import "fmt"

type teaFactory struct {

}

func (teaFactory) createBaseIngredient() {
	fmt.Println("boiling water...")
}

func (teaFactory) createMainIngredient() {
	fmt.Println("adding tea leafs...")
}

func (teaFactory) createTopper() {
	fmt.Println("adding topper...")
}
