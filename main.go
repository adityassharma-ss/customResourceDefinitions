package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MyResource struct {
	Name string `json:"name"`
	Foo  string `json:"foo"`
	Bar  int    `json:"bar"`
}

var myResources = []MyResource{}

func getMyResources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myResources)
}

func createMyResource(w http.ResponseWriter, r *http.Request) {
	var myResource MyResource
	_ = json.NewDecoder(r.Body).Decode(&myResource)
	myResources = append(myResources, myResource)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myResource)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/myresources", getMyResources).Methods("GET")
	r.HandleFunc("/myresources", createMyResource).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
