package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*go f1()
	go f2()
	fmt.Println("End of the main")
	time.Sleep(1 * time.Second)*/

	/*wg := sync.WaitGroup{} //projemızdekı join pointleri yönetmemizi sağlıyor
	wg.Add(1)              // kaç go routin'i beklediğimzii belırttik

	go func() {
		defer wg.Done()
		fmt.Println("f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()
	wg.Wait() // Blocked.wg.Wait deeik yanı usttekı go funclar bıtmeden alt satıra geçmeyecek bekle dıyoruz bı nevı

	fmt.Println("End of the main") */ // Passed time : 9.003754167s
	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("f1")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f3")
		time.Sleep(3 * time.Second)
	}()
	wg.Wait()

	fmt.Println("Passed time :", time.Now().Sub(startTime)) // geçen süre farkını hespalamak için bu sub fonksşyonunu kulllandık

}

/*func f1() {
	fmt.Println("f1")

}
func f2() {
	fmt.Println("f2")
}
*/
