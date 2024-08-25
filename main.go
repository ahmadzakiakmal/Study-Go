package main

import (
	"golang-backend/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", handler.RootRouteHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting server on port 5000...")
	
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}