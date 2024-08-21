package main

import "fmt"

func main() {

	//var numbers = []int{1, 2, 3, 4}
	/*for index := 0; index < len(numbers); index++ {
	}
	fmt.Println(numbers)*/

	/*	for index, value := range numbers { // range'in sağına hangi dizi üzerinde dolaşacaksak onu yazarız
			//fmt.Println(index, value) // index ve valueyi birlikte yazmalıyız tekini kabul etmemektedir
			fmt.Println(index, value)
		}
	*/

	/*// index'e ihtiyacımız yoksa for'un hemen sağında index yazan yere _ işareti koyamalıyız
	for _, value := range numbers { // range'in sağına hangi dizi üzerinde dolaşacaksak onu yazarız
		fmt.Println(value)
	}*/

	/* var language = "Golang"
	for _, character := range language {
		fmt.Printf("Character: %c \n", character)
	}*/

	var names = map[string]int{
		"Mustafa": 10,
		"bartu":   11,
		"bayram":  12,
	}
	for key, value := range names {
		fmt.Printf("key: %s, value: %d \n", key, value)
	}
}
