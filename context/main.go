package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")
	// context.Background() -> Boş bir context oluştuuryor.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()
	// go getProductDetailsFromExternalService(10)
	// select {
	// case productDetail := <-productChannel:
	// 	fmt.Println("Ürün detayları getirildi", productDetail)

	// case <-ctx.Done():
	// 	fmt.Println("Timeout occured, context cancelleds")
	// }
	// fmt.Println("End of the main")

	F1(ctx)
}
func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}
func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(10 * time.Second)

	productChannel <- Product{
		Id:   productId,
		Name: "Çamaşır makinesi",
	}
}
