package handler

import (
	"golang-backend/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func RootRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "layout.html"), path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{} {
		"title": "Golang Web",
		"content": "Index Page",
	} 
	
	executeErr := tmpl.Execute(w, data)
	if executeErr != nil {
		log.Println(executeErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id) 

	// data := map[string]interface{} {
	// 	"title" : "Product",
	// 	"id": id,
	// } 

	data := entity.Product{
		Id: idNum,
		Name: "Product " ,
		Price: 5000,
		Stock: 20,
	}

	if err != nil || idNum < 0 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "layout.html"), path.Join("views", "product.html"))
	if err != nil {
		log.Println(err, " Error parsing html files")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	executeErr := tmpl.Execute(w, data)
	if executeErr != nil {
		log.Println(executeErr, " Error on data execute")
		log.Println(data)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}