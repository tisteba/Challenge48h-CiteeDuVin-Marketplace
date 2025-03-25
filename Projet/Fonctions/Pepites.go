package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataPepites struct {
}

func PepitesPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		PepitesPageGet(w, r)

	case "POST":
		PepitesPagePost(w, r)
	}
}

func PepitesPageGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/pepite.html"))

	data := DataPepites{}
	tmpl.Execute(w, data)
}

func PepitesPagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/NosPepites", http.StatusSeeOther)
	fmt.Println("")
}
