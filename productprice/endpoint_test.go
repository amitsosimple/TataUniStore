/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	h "./server/handler"
)

func TestGetProductPrice_SuccessWithDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/productPrice/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetProductPrice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	productInfo := "{\"product_id\":\"2\",\"price\":{\"range\":\"12.00-22.00\",\"min\":12,\"max\":22}}\n"
	if rr.Body.String() != productInfo {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), productInfo)
	}
}

func TestCreateProductPrice_SuccessWithDetails(t *testing.T) {
	productInfo := "{\"min\":12,\"max\":22}"
	req, err := http.NewRequest("POST", "/productPrice/{id}", bytes.NewBuffer([]byte(productInfo)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req = mux.SetURLVars(req, map[string]string{
		"id": "x",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateProductPrice)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}