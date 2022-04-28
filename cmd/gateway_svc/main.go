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

	go func() {
		var data string
		req, err := http.NewRequest("GET", "10.105.16.17/api/v1/product", nil)

		// // add authorization header to the req
		// req.Header.Add("Authorization", "Bearer "+t)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
			return
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Error while unmarshalling ", err)
			return
		}

		res <- data

	}()

	fmt.Println(res)
	j, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(j)

}
