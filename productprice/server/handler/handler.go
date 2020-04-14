/**** Amit Chatter (amitsosimple@gmail.com) ****/

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"

	"../../../productutil/config"
	serviceconfig "../config"
)

type productPrice struct {
	Product_id 	string `json:"product_id"`
	Price		config.Price `json:"price,omitempty"`
}

func CreateProductPrice(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	var price config.Price

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter fields information to update")
		return
	}
	json.Unmarshal(reqBody, &price)

	products := serviceconfig.GetProductList()

	for i, singleProduct := range products {
		if singleProduct.Product_id == productID {
			if singleProduct.Price != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Product id = %s, already have price information", productID)
				return
			}

			singleProduct.Price = &config.Price{}
			//if price.Min != singleProduct.Price.Min {
			//	singleProduct.Price.Min = price.Min
			//}
			//if price.Max != singleProduct.Price.Max {
			//	singleProduct.Price.Max = price.Max
			//}
			singleProduct.Price.Min = price.Min
			singleProduct.Price.Max = price.Max

			singleProduct.Price.Range = fmt.Sprintf("%0.2f-%0.2f", singleProduct.Price.Min, singleProduct.Price.Max)

			singleProduct.UpdatedAt.Date = time.Now().Format("2006-01-02T15:04:05.999Z")
			products = append(products[:i], singleProduct)

			error := serviceconfig.UpdateCatalogInfo(products)
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Information update failure, Please retry the operation and make sure you have proper permission")
				return
			}

			json.NewEncoder(w).Encode(singleProduct)
			return
		}
	}
	fmt.Fprintf(w, "The product with ID %v is not available for modification", productID)
}

func GetProductPrice(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]

	products := serviceconfig.GetProductList()

	for _, singleProduct := range products {
		if singleProduct.Product_id == productID {

			if singleProduct.Price == nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Price for the ProductId=%s is not available", productID)
				return
			}

			productPrice := &productPrice{
				Product_id: singleProduct.Product_id,
				Price: *(singleProduct.Price),
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(productPrice)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "The product with ID %v is not available for modification", productID)
}
