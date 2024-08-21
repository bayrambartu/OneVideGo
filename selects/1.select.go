package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	var data1 string
	var data2 string

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "World"
	}()

	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 = <-channel1:
			fmt.Println("reseived data from channel1: ", data1)

		case data2 = <-channel2:
			fmt.Println("received data form channel2:", data2)
		default: // her iki channel'a bilgi gelmezse default değeri yazıdracaktır.
			fmt.Println("No data came..")
			time.Sleep(1 * time.Second)
		}
	}

}
