package Fonctions

import (
	"fmt"
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

type DataEncheres struct {
	WineList  []Wine
	IsConnect bool
	User      db.DB
}

func EncheresPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		EncheresGet(w, r)

	case "POST":
		EncheresPost(w, r)
	}
}

func EncheresGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/encheres.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}

	wines, err := LoadWines()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier JSON :", err)
		return
	}

	var Instance []Wine
	for _, wine := range wines {
		if wine.Price == 0 {
			Instance = append(Instance, wine)
		}
	}

	data := DataEncheres{
		WineList:  Instance,
		IsConnect: IsConnectStock,
		User:      db.GetUser(cookie.Value),
	}

	tmpl.Execute(w, data)
}

func EncheresPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Encheres", http.StatusSeeOther)
}
