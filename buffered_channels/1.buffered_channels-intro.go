package main

import "fmt"

func main() {

	myChannel := make(chan int, 2)
	myChannel <- 1
	myChannel <- 2
	fmt.Println(<-myChannel) // myChannel'dan bilgi okuma işlemi.

	// burda aslında myChanlle'a ait 1 tane veri okuduk yanı channel'ımızın size'ı 1 oldu ondan dolayı bir tane daha veri basabılırız.
	myChannel <- 3
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

/*Bu programın açıklaması:

1.myChannel adında 2 kapasiteli bir buffer'lı kanal oluşturulur.
2.Kanala 1 ve 2 değerleri gönderilir.
3.fmt.Println(<-myChannel) ifadesiyle, kanaldan bir veri alınır ve yazdırılır. Bu, kanalın boyutunu 1 yapar.
4.Kanala tekrar 3 değeri gönderilir. Şu anda kanalın kapasitesi 1 ve dolu 1 yeri olduğu için, yeni veri gönderme işlemi başarılı olur.
5.Kanalın kalan verileri sırayla alınır ve yazdırılır.

Programda, kanalın her zaman yeterli kapasitesi olduğu için deadlock hatası oluşmaz.*/
