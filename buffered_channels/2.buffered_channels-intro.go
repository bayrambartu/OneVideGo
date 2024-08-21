package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan string, 2)

	go func() {
		data1 := <-message
		fmt.Println("First:", data1)
		data2 := <-message
		fmt.Println("Second:", data2)
		message <- "Hello"
		message <- "World"
		message <- "World2"
		time.Sleep(1 * time.Second)
	}()
	fmt.Println("End of the main ")
}

/* Aslınfa goroutine çalıştı ancak main Go routine bittiği için çıktı olarak
End of the main'i görürüz */
