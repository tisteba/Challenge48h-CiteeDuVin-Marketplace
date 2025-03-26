package Fonctions

import (
	"html/template"
	"net/http"
)

type DataEncheres struct {
}

func EncheresPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		EncheresGet(w, r)

	case "POST":
		EncheresPost(w, r)
	}
}

func EncheresGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/encheres.html"))

	data := DataEncheres{}
	tmpl.Execute(w, data)
}

func EncheresPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Encheres", http.StatusSeeOther)
}
