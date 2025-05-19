package main

import (
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, _ := DB.Query("SELECT id, name, description, price FROM products")
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		products = append(products, p)
	}
	templates.ExecuteTemplate(w, "index.html", products)
}

func newProductHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "form.html", nil)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		DB.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", name, description, price)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := DB.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?", id)

	var p Product
	row.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
	templates.ExecuteTemplate(w, "form.html", p)
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

	DB.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?", name, description, price, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	DB.Exec("DELETE FROM products WHERE id = ?", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
