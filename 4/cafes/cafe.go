package cafes

import (
	"3/cafes/coffeeHouse/products"
	"fmt"
)

type ICafe interface {
	Menu()
	NewOrder()
	Order()
}

type order struct {
	Products []products.IProduct
}

type Cafe struct {
	Orders []order
}

func (cafe Cafe) Order() (string, error) {

	var order = "***Order***\n"

	if len(cafe.Orders) == 0 {
		return "", fmt.Errorf("no orders")
	}

	lastOrder := cafe.Orders[len(cafe.Orders)-1]

	if len(lastOrder.Products) == 0 {
		return "", fmt.Errorf("no products in order")
	}

	for _, product := range lastOrder.Products {
		order += fmt.Sprintf("%v %v$\n", product.Description(), product.Cost())
	}

	order += "***********"

	return order, nil
}

func (cafe *Cafe) NewOrder() {
	cafe.Orders = append(cafe.Orders, order{})
}
