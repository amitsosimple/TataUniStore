/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"fmt"
)

func testAggregator() {
	fmt.Println("Starting the application...")
	ReturnProductInfoByProductId(2)
	fmt.Println()
	GetPriceInfoByProductId(2)
	fmt.Println()
	CreateNewProduct()
	fmt.Println()
	GetAllProductsDetails()
	fmt.Println()
	CreateNewPrice(2)
	fmt.Println()
	UpdateProductInfoByProductId(1)
	fmt.Println()
}

func main() {
	testAggregator()
}
