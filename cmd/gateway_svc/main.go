package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/hello", sayHello).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world gateway")
}
