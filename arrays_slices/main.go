package main

import "fmt"

func main() {

	//fixed
	/* var names [3]string
	names[0] = "Mustafa"
	names[1] = "bartu"
	names[2] = "bayram"
	fmt.Println(names) */

	/* var names = [4]string{"Mustafa", "bartu", "bayram", "serhat"}

	fmt.Println(names[0:3])*/

	var names = []string{"ali", "kenan", "serhat", "tolga", "can"}
	names = append(names, "bartu")
	fmt.Println(names)
}
