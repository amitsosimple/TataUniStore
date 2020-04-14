package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"

	"../../../productutil"
	"../../../productutil/config"
	u "../../../productutil/log"
	serviceconfig "../config"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer u.Exit(u.Enter())
	var newProduct serviceconfig.ProductInfo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.ErrorLogger.Println("Bad Request for Product Creation")
		fmt.Fprintf(w, "Bad Request for Product Creation")
		return
	}

	json.Unmarshal(reqBody, &newProduct)

	if len(newProduct.Product_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		u.ErrorLogger.Println("Bad Request for Product Creation, Please provide Product_Id")
		fmt.Fprintf(w, "Bad Request for Product Creation, Please provide Product_Id")
		return
	}

	products := serviceconfig.GetProductList()

	for _, singleProduct := range products {
		if singleProduct.Product_id == newProduct.Product_id {
			w.WriteHeader(http.StatusBadRequest)
			u.ErrorLogger.Printf("Bad Request for Product Creation, Product with Product_Id=%s is already present\n", newProduct.Product_id)
			fmt.Fprintf(w, "Bad Request for Product Creation, Product with Product_Id=%s is already present", newProduct.Product_id)
			return
		}
	}

	newProduct.PublishedAt.Date = time.Now().Format("2006-01-02T15:04:05.999Z")
	newProduct.CreatedAt.Date = time.Now().Format("2006-01-02T15:04:05.999Z")
	newProduct.UpdatedAt.Date = time.Now().Format("2006-01-02T15:04:05.999Z")

	products = append(products, serviceconfig.TranslateToProduct(newProduct))

	error := serviceconfig.UpdateCatalogInfo(products)
	if error != nil {
		u.ErrorLogger.Println("Information update failure, Please retry the operation and make sure you have proper permission")
		fmt.Fprintf(w, "Information update failure, Please retry the operation and make sure you have proper permission")
		return
	}
	u.GeneralLogger.Printf("Product Creation successful with ProductId=%s", newProduct.Product_id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(products)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	defer u.Exit(u.Enter())
	productID := mux.Vars(r)["id"]

	products := serviceconfig.GetProductList()

	for _, singleProduct := range products {
		if singleProduct.Product_id == productID {
			u.GeneralLogger.Printf("Successfully fetched and displyed the catalog information of productID=%s\n", productID)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(serviceconfig.TranslateToProductInfo(singleProduct))
			return
		}
	}
	u.ErrorLogger.Printf("The product with ID %v is not available to display\n", productID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The product with ID %v is not available to display", productID)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	defer u.Exit(u.Enter())
	products := serviceconfig.GetProductList()

	if(len(products) == 0) {
		u.GeneralLogger.Println("Product catalog is empty, No Product to display")
		fmt.Fprint(w, "Product catalog is empty, No Product to display")
		return
	}

	u.GeneralLogger.Println("Product List displayed successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	defer u.Exit(u.Enter())
	products := serviceconfig.GetProductList()
	if len(products) == 0 {
		u.GeneralLogger.Println("Product catalog is empty, No Product to display")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Product catalog is empty, No Product to Update")
		return
	}

	productID := mux.Vars(r)["id"]
	var updatedProduct config.Product

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.GeneralLogger.Println("Kindly enter fields information to update")
		fmt.Fprintf(w, "Kindly enter fields information to update")
		return
	}
	json.Unmarshal(reqBody, &updatedProduct)

	for i, singleProduct := range productutil.Products {
		if singleProduct.Product_id == productID {
			if  singleProduct.Product_id != updatedProduct.Product_id {
				u.ErrorLogger.Printf("Data Mismatched, Query param and body data is not match, Param Id=%s where the Payload Id=%s\n", productID, updatedProduct.Product_id)
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w,"Data Mismatched, Query param and body data is not match, Param Id=%s where the Payload Id=%s", productID, updatedProduct.Product_id)
				return
			}
			if updatedProduct.SellerId != singleProduct.SellerId {singleProduct.SellerId = updatedProduct.SellerId}
			if updatedProduct.Title != singleProduct.Title {singleProduct.Title = updatedProduct.Title}
			if updatedProduct.Manufacturer != singleProduct.Manufacturer {singleProduct.Manufacturer = updatedProduct.Manufacturer}
			if updatedProduct.IsLowQuantity != singleProduct.IsLowQuantity {singleProduct.IsLowQuantity = updatedProduct.IsLowQuantity}
			if updatedProduct.IsSoldOut != singleProduct.IsSoldOut {singleProduct.IsSoldOut = updatedProduct.IsSoldOut}
			if updatedProduct.IsBackorder != singleProduct.IsBackorder {singleProduct.IsBackorder = updatedProduct.IsBackorder}
			if updatedProduct.RequiresShipping != singleProduct.RequiresShipping {singleProduct.RequiresShipping = updatedProduct.RequiresShipping}
			if updatedProduct.IsVisible != singleProduct.IsVisible {singleProduct.IsVisible = updatedProduct.IsVisible}
			if updatedProduct.Metafields != nil {
				for _, NewMetaField := range updatedProduct.Metafields {
					noMatch := true
					for i, MetaField := range singleProduct.Metafields {
						if MetaField.Key == NewMetaField.Key {
							singleProduct.Metafields[i].Value = NewMetaField.Value
							noMatch = false
						}
					}
					if(noMatch) {
						singleProduct.Metafields = append(singleProduct.Metafields, NewMetaField)
					}
				}
			}
			singleProduct.Workflow.Status = "updated"
			//if updatedProduct.Price != nil {
			//	if updatedProduct.Price.Min != singleProduct.Price.Min {
			//		singleProduct.Price.Min = updatedProduct.Price.Min
			//	}
			//	if updatedProduct.Price.Max != singleProduct.Price.Max {
			//		singleProduct.Price.Max = updatedProduct.Price.Max
			//	}
			//	singleProduct.Price.Range = fmt.Sprintf("%0.2f-%0.2f", singleProduct.Price.Min, singleProduct.Price.Max)
			//}
			singleProduct.UpdatedAt.Date = time.Now().Format("2006-01-02T15:04:05.999Z")
			products = append(products[:i], singleProduct)

			error := serviceconfig.UpdateCatalogInfo(products)
			if error != nil {
				fmt.Fprintf(w, "Information update failure, Please retry the operation and make sure you have proper permission")
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(singleProduct)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "The product with ID %v is not available for modification", productID)
}
