package main

import "fmt"

func main() {
	var condition = true
	defer func() {
		fmt.Println("cleanup worked..")
	}()

	if condition {
		panic("An error occurred") // uygulama bir anda tamamen sonlanması gerekiyorsa bu durumlarda kullanılan fonksiyondur.
	} // defer'ler de tamamen sonlanma ılmadan hemen önce kullanılıdıgı için defer çalışıtırlır sonrasında panic içindeki yazı yazdırılır.

	// cleanup() burada kaldırıldı çünkü defer zaten fonksiyonu çalıştıracak
}
