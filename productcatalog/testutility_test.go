/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"./server/config"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func createProductInfo() {
	os.Remove(config.GetProductFileName())
	config.InitCatalogClient()
	productInfo := "[{\"product_id\":\"2\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true,\"publishedAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"createdAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"updatedAt\":{\"$date\":\"2020-04-13T22:27:34.003Z\"},\"workflow\":{\"status\":\"new\"},\"price\":{\"range\":\"12.00-22.00\",\"min\":12,\"max\":22}}]"
	ioutil.WriteFile(config.GetProductFileName(), []byte(productInfo), 0644)
}

func cleanProductInfo() {
	os.Remove(config.GetProductFileName())
}

func setup() {
	createProductInfo()
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	cleanProductInfo()
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}
