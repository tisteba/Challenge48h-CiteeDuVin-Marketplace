package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataHome struct {
	WineList []Wine
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

	wines, err := LoadWines()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier JSON :", err)
		return
	}

	data := DataHome{
		WineList: wines,
	}
	tmpl.Execute(w, data)
}

func HomePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
	fmt.Println("")
}
