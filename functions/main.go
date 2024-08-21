package main

import "fmt"

func main() {
	/*total := add(10, 20)
	fmt.Println(total) */
	/*total, difference, multiply := calculation(10, 20)
	fmt.Println(total, difference, multiply)*/
	/*var numbers = [5]int{10, 20, 2, 3, 5}
	fmt.Println(sum(numbers[:]))*/
	/*fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	*/
	fmt.Println(sum(1, 2, 3, 4, 45, 10, 10, 10, 10, 5))
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*func sum(x int, y int, z int) int {
	return x + y + z

}*/

/*func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}*/

func calculation(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func add(x int, y int) int {
	return x + y
}
