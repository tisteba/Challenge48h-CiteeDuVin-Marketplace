package Fonctions

import (
	"html/template"
	"net/http"
)

func PanierPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("../Front/HTML/panier.html"))
		tmpl.Execute(w, nil)
	}
}
