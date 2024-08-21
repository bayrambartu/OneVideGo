package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "Bayram bartu kurnaz", age: 21,
		adress: Adress{"istanbul", "Üsküdar"}}

	customer1.ToString()
	customer1.changeName("Swift")
	customer1.ToString()

	//var customer2 = Customer{id: 2, name: "Mustafa murat coşkun", age: 55}

	/*customer1.age = 31
	fmt.Println(customer1)
	fmt.Println(customer2)*/
}
func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct {
	id     int64
	name   string
	age    int
	adress Adress
}
type Adress struct {
	city     string
	district string
}

func (customer *Customer) ToString() { // customer struct'ının fonksiyonu halıne getirdik
	fmt.Printf("Name: %s ,Age: %d \n", customer.name, customer.age)
}
