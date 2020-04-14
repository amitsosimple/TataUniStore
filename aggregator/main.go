/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"fmt"

	u "../productutil/log"
)

func testAggregator() {
	defer u.Exit(u.Enter())
	u.GeneralLogger.Println("Starting the application...")
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
