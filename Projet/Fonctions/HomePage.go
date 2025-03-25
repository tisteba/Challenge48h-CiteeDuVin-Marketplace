package Fonctions

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
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

	var wine9 []Wine
	for i := 0; i < 9; i++ {
		rand.Seed(time.Now().UnixNano()) // Initialisation avec le temps actuel pour éviter la répétition
		nombre := rand.Intn(10000)

		wine9 = append(wine9, wines[nombre])
	}

	data := DataHome{
		WineList: wine9,
	}
	tmpl.Execute(w, data)
}

func HomePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
	fmt.Println("")
}

func RedirectBasePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
}
