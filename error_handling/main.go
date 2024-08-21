package main

import (
	"fmt"
)

func main() {
	productservice := ProductService{}
	err := productservice.Add(Product{id: 1, name: "Çamaşır", price: 1000})

	if err != nil {
		fmt.Println(err)
	}

	/*result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}*/
	/*fileData, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}*/

}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

func (productservice *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Ürün ismi boş olamaz", code: "1001"}
	}
	if product.price < 0 {
		return ValidationError{text: "Ürün fiyatı 0'dan büyük olmalıdır", code: "1002"}

	}
	fmt.Println("Ürün eklendi")
	return nil
}

type ValidationError struct {
	text string
	code string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata: %s,Kod: %s", validationError.text, validationError.code)
}

/*func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("payda sıfır olamaz")
	} else {
		return x / y, nil
	}
}*/
