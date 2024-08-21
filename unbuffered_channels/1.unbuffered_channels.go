package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	done := make(chan bool)
	go func() {
		message := <-myChannel // message myChanller'den bi bilgi okumaktadır.
		fmt.Println(message)
		done <- true
	}()
	go func() {
		myChannel <- "hello World"
	}()
	<-done
	fmt.Println("End of the main")
}
