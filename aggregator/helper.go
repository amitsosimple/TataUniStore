/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func printResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)

		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, data, "", "\t")
		if error != nil {
			fmt.Printf(string(data))
		return
		}
		fmt.Println(&prettyJSON)
	}
}

func ReturnProductInfoByProductId(id int) {
	endpoint := fmt.Sprintf("http://localhost:8081/products/%d",id)
	response, err := http.Get(endpoint)
	printResponse(response, err)
}

func GetAllProductsDetails() {
	endpoint := fmt.Sprintf("http://localhost:8081/products")
	response, err := http.Get(endpoint)
	printResponse(response, err)
}

func GetPriceInfoByProductId(id int){
	endpoint := fmt.Sprintf("http://localhost:8082/productPrice/%d",id)
	response, err := http.Get(endpoint)
	printResponse(response, err)
}

func CreateNewProduct() {
	payload := "{\"product_id\":\"2\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true,\"publishedAt\":{\"$date\":\"2020-02-12T08:05:39.743Z\"},\"createdAt\":{\"$date\":\"2010-08-23T05:53:16.134Z\"},\"updatedAt\":{\"$date\":\"2019-08-23T05:53:16.134Z\"},\"workflow\":{\"status\":\"new\"},\"price\":{\"min\":1,\"max\":2}}"

	var buf bytes.Buffer
	_, _ = buf.Write([]byte(payload))

	response, err := http.Post("http://localhost:8081/product", "application/json", &buf)
	printResponse(response, err)
}

func CreateNewPrice(id int) {
	endpoint := fmt.Sprintf("http://localhost:8082/productPrice/%d",id)

	payload := "{\"min\":12,\"max\":22}"
	var buf bytes.Buffer
	_, _ = buf.Write([]byte(payload))

	response, err := http.Post(endpoint, "application/json", &buf)
	printResponse(response, err)
}

func UpdateProductInfoByProductId(id int){
	endpoint := fmt.Sprintf("http://localhost:8081/products/%d",id)
	payload := "{\"product_id\":\"2\",\"sellerId\":\"3\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true,\"publishedAt\":{\"$date\":\"2020-02-12T08:05:39.743Z\"},\"createdAt\":{\"$date\":\"2010-08-23T05:53:16.134Z\"},\"updatedAt\":{\"$date\":\"2019-08-23T05:53:16.134Z\"},\"workflow\":{\"status\":\"new\"},\"price\":{\"min\":1,\"max\":2}}"

	var buf bytes.Buffer
	_, _ = buf.Write([]byte(payload))

	// initialize http client
	client := &http.Client{}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	printResponse(resp, err)
}