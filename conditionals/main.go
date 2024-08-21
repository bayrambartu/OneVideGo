package main

import "fmt"

func main() {

	/*var age = 17

	if age >= 18 {
		fmt.Println("you can vote!")
	} else {
		fmt.Println("you can not vote!")
	}*/
	a := 10
	b := 100
	c := 30

	if a >= b && a >= c {
		fmt.Println("Greatest variable is a!")
	} else if b >= a && b >= c {
		fmt.Println("Greatest variable is b!")

	} else {
		fmt.Println("Greatest variable is c!")

	}

}
