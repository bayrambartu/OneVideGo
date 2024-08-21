package main

import (
	"fmt"
)

func main() {

	// bu örnekte main go routine bir tane unbuffered channel'a veri basacak
	myChannel := make(chan int)

	go func() {
		data := <-myChannel
		fmt.Println("First Go routine took data: ", data)

	}()
	go func() {
		data := <-myChannel
		fmt.Println("Second Go routine took data: ", data)

	}()

	myChannel <- 10
	fmt.Println("End of the main")
}
