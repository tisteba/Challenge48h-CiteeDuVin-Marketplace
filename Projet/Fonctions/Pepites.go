package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
)

type DataPepites struct {
	WineList []Wine
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

	wines, err := LoadWines()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier JSON :", err)
		return
	}

	SortWinesByPoints(wines)

	var wine12 []Wine
	for i := len(wines) - 1; i > len(wines)-13; i-- {
		wine12 = append(wine12, wines[i])
	}

	data := DataPepites{
		WineList: wine12,
	}
	tmpl.Execute(w, data)
}

func PepitesPagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/NosPepites", http.StatusSeeOther)
	fmt.Println("")
}
