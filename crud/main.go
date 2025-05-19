package main

import (
	"net/http"
)

func main() {
	initDatabase()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newProductHandler)
	http.HandleFunc("/create", createProductHandler)
	http.HandleFunc("/edit", editProductHandler)
	http.HandleFunc("/update", updateProductHandler)
	http.HandleFunc("/delete", deleteProductHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
