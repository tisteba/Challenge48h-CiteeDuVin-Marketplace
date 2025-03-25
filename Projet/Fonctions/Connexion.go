package Fonctions

import (
	"html/template"
	"net/http"
)

var StockMessageErreurConnexion string

type DataConnexion struct {
	MessageErreur string
}

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ConnexionGet(w, r)

	case "POST":
		ConnexionPost(w, r)
	}
}

func ConnexionGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/connexion.html"))

	data := DataConnexion{
		MessageErreur: StockMessageErreurConnexion,
	}
	tmpl.Execute(w, data)
}

func ConnexionPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
}
