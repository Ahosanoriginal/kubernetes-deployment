package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name string
}

var products = []Product{
	{Name: "toto"},
	{Name: "titi"},
	{Name: "hello"},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/product", getProducts).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":8001",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(products)
	w.Write(j)
}
