package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataHome struct {
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		HomePageGet(w, r)

	case "POST":
		HomePagePost(w, r)
	}
}

func HomePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/homepage.html"))

	data := DataHome{}
	tmpl.Execute(w, data)
}

func HomePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
	fmt.Println("")
}
