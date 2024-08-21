package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("sent data", i)
			time.Sleep(1 * time.Second)

		}
		close(myChannel) // eğer channlerimizi kapatmazwak deadlock'a girer .Artık channel'a bilgi basmayacağız basmadığımız zaman da isOpen =false olur ve döngümüzden break'la çıkmış olacağız.
	}()
	for {
		data, isOpen := <-myChannel
		if isOpen == false {
			break
		}
		fmt.Println("Received data: ", data)
	}
}
