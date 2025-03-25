package Fonctions

import (
	"fmt"
	"html/template"
	db "marketplace/DataBase"
	"math/rand"
	"net/http"
	"time"
)

var IsConnectStock bool

type DataHome struct {
	WineList  []Wine
	IsConnect bool
	User      db.DB
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
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 9; i++ {
		nombre := rand.Intn(10000)

		wine9 = append(wine9, wines[nombre])
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}

	data := DataHome{
		WineList:  wine9,
		IsConnect: IsConnectStock,
		User:      db.GetUser(cookie.Value),
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

func SetUpConnect() {
	IsConnectStock = true
}
