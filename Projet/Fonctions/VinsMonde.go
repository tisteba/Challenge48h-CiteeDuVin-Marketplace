package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataMonde struct {
}

func VinsMondePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		VinsMondePageGet(w, r)

	case "POST":
		VinsMondePagePost(w, r)
	}
}

func VinsMondePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/vinsmonde.html"))

	data := DataPepites{}
	tmpl.Execute(w, data)
}

func VinsMondePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/VinsDuMonde", http.StatusSeeOther)
	fmt.Println("")
}
