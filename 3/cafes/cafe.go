package cafes

import (
	"3/cafes/coffeeHouse/products"
	"fmt"
)

const (
	darkroast = iota + 1
	houseblend
	decaf
	espresso
)

type ICafe interface {
	PrintMenu()
	newOrder()
	PrintOrder()
}

type order struct {
	products []products.IProduct
}

type Cafe struct {
	orders []order
}

func (cafe Cafe) Order() (string, error) {

	var order = "***Order***\n"

	if len(cafe.orders) == 0 {
		return "", fmt.Errorf("no orders")
	}

	lastOrder := cafe.orders[len(cafe.orders)-1]

	if len(lastOrder.products) == 0 {
		return "", fmt.Errorf("no products in order")
	}

	for _, product := range lastOrder.products {
		order += fmt.Sprintf("%v %v $\n", product.Description(), product.Cost())
	}

	order += "***********"

	return order, nil
}

func (cafe *Cafe) NewOrder() {
	cafe.orders = append(cafe.orders, order{})
}
