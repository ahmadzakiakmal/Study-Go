package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", rootRouteHandler)

	log.Println("Starting server on port 5000...")
	
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

func rootRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
}
	w.Write([]byte("Hello world! I'm learning Golang Backend!"))
}