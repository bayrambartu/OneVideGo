package main

import "fmt"

func main() {
	/*var names map[string]int
	names = make(map[string]int, 0) // 0'ın anlamı map'in içinde herhangi bir değerinin olmadığı anlamındadır

	names["Bartu"] = 1 // Bartuya karşılık gelen değerim 1'dir
	names["Bayram"] = 2
	names["Kurnaz"] = 3
	// toplamda 3 key oluşturduk ve bu key'lere karşılık gelen değerler bu şekilde

	fmt.Println(names["Melike"]) */

	names := map[string]int{
		"Bartu":  1,
		"Bayram": 2,
		"Kurnaz": 3,
	}

	//daha sonradan şstemedğimiz key'ler olursa silebilriiz-->
	delete(names, "Bartu") //

	fmt.Println(names)
}
