package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}
type Electronic struct {
	desi int
}
type Flower struct {
	desi int
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}
func (book *Book) ShippingCost() int {
	return 5 + book.desi*2

}
func (flower *Flower) ShippingCost() int {
	return 12 + flower.desi*3
}
func main() {
	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 20},
		&Flower{desi: 10},
	}
	fmt.Println(calculateTotalShippingCost(products))

	/*var product IShippable

	product = &Book{desi: 10}
	fmt.Println(product.ShippingCost())

	product = &Electronic{desi: 10}

	fmt.Println(product.ShippingCost())*/

	/*var electronic *Electronic
	var book *Book
	book1 := &Book{desi: 10}
	book2 := &Book{desi: 20}

	fmt.Println(book1.ShippingCost())
	fmt.Println(book2.ShippingCost())*/
	//books := []Book{Book{desi: 10}, Book{desi: 20}}
	//fmt.Println(calculateTotalShippingCost(books))
	/*electronic1 := &Electronic{desi: 20}
	fmt.Println(electronic1.ShippingCost())*/

	/*electronics := []Electronic{{desi: 10}, {desi: 20}}
	calculateTotalShippingCostOfElectronics(electronics)*/
}
func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total

}

/*func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, books := range books {
		total += books.ShippingCost()
	}
	return total

}
func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronics := range electronics {
		total += electronics.ShippingCost()
	}
	return total

}
*/
