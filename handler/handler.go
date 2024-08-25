package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func RootRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
}
	w.Write([]byte("Hello world! I'm learning Golang Backend!"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id) 

	if err != nil || idNum < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product id: %d", idNum)
}