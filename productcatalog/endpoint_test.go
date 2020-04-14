/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "./server/handler"
)

func TestGetAllProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetAllProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	productInfo := "[{\"product_id\":\"2\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true,\"publishedAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"createdAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"updatedAt\":{\"$date\":\"2020-04-13T22:27:34.003Z\"},\"workflow\":{\"status\":\"new\"},\"price\":{\"range\":\"12.00-22.00\",\"min\":12,\"max\":22}}]\n"
	if rr.Body.String() != productInfo {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), productInfo)
	}
}

func TestGetOneProduct_SuccessWithDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetOneProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	productInfo := "{\"product_id\":\"2\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true,\"publishedAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"createdAt\":{\"$date\":\"2020-04-13T22:27:34.002Z\"},\"updatedAt\":{\"$date\":\"2020-04-13T22:27:34.003Z\"},\"workflow\":{\"status\":\"new\"}}\n"
	if rr.Body.String() != productInfo {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), productInfo)
	}
}

func TestGetOneProduct_SuccessWithNoDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetOneProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	productInfo := fmt.Sprintf("The product with ID 1 is not available to display")
	if rr.Body.String() != productInfo {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), productInfo)
	}
}

func TestCreateProduct_SuccessWithDetails(t *testing.T) {
	productInfo := "{\"product_id\":\"1\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true}"
	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer([]byte(productInfo)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestCreateProduct_BadRequestAlreadyPresent(t *testing.T) {
	productInfo := "{\"product_id\":\"1\",\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true}"
	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer([]byte(productInfo)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	failure := fmt.Sprintf("Bad Request for Product Creation, Product with Product_Id=1 is already present")
	if rr.Body.String() != failure {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), failure)
	}
}

func TestCreateProduct_BadRequestNoProductId(t *testing.T) {
	productInfo := "{\"sellerId\":\"2\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true}"
	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer([]byte(productInfo)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	failure := fmt.Sprintf("Bad Request for Product Creation, Please provide Product_Id")
	if rr.Body.String() != failure {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), failure)
	}
}

func TestUpdateProduct_SuccessWithDetails(t *testing.T) {
	defer createProductInfo()
	productInfo := "{\"product_id\":\"2\",\"sellerId\":\"1\",\"title\":\"1\",\"manufacturer\":\"1\",\"isLowQuantity\":false,\"isSoldOut\":false,\"isBackorder\":false,\"metafields\":[{\"key\":\"Capacity\",\"value\":\"\"},{\"key\":\"Capacity1\",\"value\":\"\"}],\"requiresShipping\":true,\"isVisible\":true}"
	req, err := http.NewRequest("PUT", "/product", bytes.NewBuffer([]byte(productInfo)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if strings.Contains(rr.Body.String(), productInfo) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), productInfo)
	}
}