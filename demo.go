package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

import . "./controller"

func main() {
	fmt.Println("Servidor no ar.")

	router := mux.NewRouter()
	router.HandleFunc("/bid", CreateBid).Methods("POST")
	router.HandleFunc("/stats", GetStats).Methods("GET")

	LoadData()

	log.Fatal(http.ListenAndServe(":8000", router))
}
