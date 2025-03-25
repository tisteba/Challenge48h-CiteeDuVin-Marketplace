package Fonctions

import (
	"fmt"
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

type DataEnsemble struct {
	WineList  []Wine
	IsConnect bool
	User      db.DB
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

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}

	data := DataEnsemble{
		WineList:  wines,
		IsConnect: true,
		User:      db.GetUser(cookie.Value),
	}
	tmpl.Execute(w, data)
}

func EnsembleVinsPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/TousTousLesVins", http.StatusSeeOther)
}
