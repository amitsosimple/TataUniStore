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
	u.GeneralLogger.Println("HomePage of service ProductPrice Server hit")
	fmt.Fprintf(w, "Welcome home!")
}

func Serve() {
	defer u.Exit(u.Enter())
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/productPrice/{id}", handler.CreateProductPrice).Methods("POST")
	router.HandleFunc("/productPrice/{id}", handler.GetProductPrice).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}
