/**** Amit Chatter (amitsosimple@gmail.com) ****/

package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	u "../../productutil/log"
	"./handler"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	u.GeneralLogger.Println("HomePage of service ProductionCatalog Server hit")
	fmt.Fprintf(w, "Welcome home!")
}

func Serve() {
	defer u.Exit(u.Enter())
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetOneProduct).Methods("GET")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}
