package main

import (
	"fmt"
)

func main() {
	/*var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = 0.37
	isInStock = false

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))*/

	/*var productName = "Golang dersleri"
	fmt.Println(productName, reflect.TypeOf(productName))*/

	/*productName := "Golang"
	fmt.Println(productName)*/

	/* var quantity int64 = 10
	fmt.Println(quantity, reflect.TypeOf(quantity)) */

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = 0.37
	isInStock = false

	//fmt.Println("product Name :", productName, "Quantity:", quantity, "Discount :", discount, "isInStock:", isInStock)
	//fmt.Printf("product name: %s, Quantity: %b, Disscount:%f, is in stcok: %t\n", productName, quantity, discount, isInStock)
	fmt.Printf("product name: %v, Quantity: %v, Disscount:%v, is in stcok: %v\n", productName, quantity, discount, isInStock)

}
