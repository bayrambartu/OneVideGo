package main

import "fmt"

func main() {
	/*	var a int
		var p *int
		a = 10
		p = &a
		fmt.Println(a)
		//pointerin göstedği adrese git ve 20 yap :
		*p = 20
		fmt.Println(a)*/

	/* var a = 10
	var b int
	var p *int

	b = a

	p = &a

	*p = 20
	fmt.Println(a, b)
	fmt.Println(&a)
	fmt.Print(&b)*/
	/*var a int = 10
	fmt.Println(a)
	add12pointer(&a)
	fmt.Println(a)*/

	// var numbers = []int{1, 2, 3}
	// fmt.Println(numbers)
	// changeValue(numbers)
	// fmt.Println(numbers)

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Print(j) // see the new vaşues of j
}

/*
func changeValue(numbers []int) { // arrayler ve slicesler referans tipler oldugu için bunlar zaten fonksiyonşara tutgu degerın pointer yanı referans degerını yollarlar
	numbers[0] = 1000 // numbers'ın 0. adresine gidip o adresteki değri değiştiriyor

}

func add12(x int) {
	x += 12

}

func add12pointer(x *int) {
	*x += 12

} */
