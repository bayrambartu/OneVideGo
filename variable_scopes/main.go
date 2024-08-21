package main

import "fmt"

var x = 10

func main() {
	// block scope
	/*var condition = true

	if condition {
		//block
		var x = 10
		fmt.Println(x)
	}*/
	/*var condition = true
	if condition {
		fmt.Println(condition)
	}
	fmt.Println(condition) // fonksiyon bitmediği için tnaımlıdır.*/
	fmt.Println(x)
	test()
	fmt.Println(x)
}
func test() { // function scope
	x = 30
	fmt.Println(x)
}
