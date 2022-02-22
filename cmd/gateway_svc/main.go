package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1", sayHello).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	res := make(chan string)
	go requestProduct(res)
	j, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(j)
}

func requestProduct(res chan string) {
	resp, err := http.Get("10.101.209.249/api/v1/product")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	res <- sb
}
