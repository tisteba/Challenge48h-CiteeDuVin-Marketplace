package Fonctions

import (
	"html/template"
	"net/http"
)

type DataRdv struct {
}

func RdvPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		RdvGet(w, r)

	case "POST":
		RdvPost(w, r)
	}
}

func RdvGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/rdv.html"))

	data := DataQuiz{}
	tmpl.Execute(w, data)
}

func RdvPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/PriseRdv", http.StatusSeeOther)
}
