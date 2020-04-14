/**** Amit Chatter (amitsosimple@gmail.com) ****/

package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"./handler"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func Serve() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetOneProduct).Methods("GET")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}
