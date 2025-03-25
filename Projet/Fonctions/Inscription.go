package Fonctions

import (
	"html/template"
	"net/http"
)

type DataInscrition struct {
}

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		InscriptionGet(w, r)

	case "POST":
		InscriptionPost(w, r)
	}
}

func InscriptionGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/inscription.html"))

	data := DataConnexion{}
	tmpl.Execute(w, data)
}

func InscriptionPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Inscription", http.StatusSeeOther)
}
