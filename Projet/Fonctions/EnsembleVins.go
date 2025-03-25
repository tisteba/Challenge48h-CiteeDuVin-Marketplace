package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataEnsemble struct {
	WineList []Wine
}

func EnsembleVinsPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		EnsembleVinsGet(w, r)

	case "POST":
		EnsembleVinsPost(w, r)
	}
}

func EnsembleVinsGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/ensemble.html"))

	wines, err := LoadWines()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier JSON :", err)
		return
	}

	data := DataEnsemble{
		WineList: wines,
	}
	tmpl.Execute(w, data)
}

func EnsembleVinsPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/TousTousLesVins", http.StatusSeeOther)
}

func RedirectBasePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/TousLesVins", http.StatusSeeOther)
}
