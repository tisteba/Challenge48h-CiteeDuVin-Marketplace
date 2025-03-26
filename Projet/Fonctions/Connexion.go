package Fonctions

import (
	"fmt"
	"html/template"
	db "marketplace/DataBase"
	"net/http"
	"strconv"
	"time"
)

var StockMessageErreurConnexion string

type DataConnexion struct {
	MessageErreur string
	IsConnect     bool
	User          db.DB
}

func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ConnexionGet(w, r)

	case "POST":
		ConnexionPost(w, r)
	}
}

func ConnexionGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/connexion.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
		return
	}

	data := DataConnexion{
		MessageErreur: StockMessageErreurConnexion,
		IsConnect:     IsConnectStock,
		User:          db.GetUser(cookie.Value),
	}

	StockMessageErreurConnexion = ""
	tmpl.Execute(w, data)
}

func ConnexionPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("Erreur ParseForm() : %v\n", err)
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	mdp := r.Form.Get("password")

	tabUser := db.GetDB()

	for _, user := range tabUser {
		if user.Email == email && user.Mdp == mdp {
			cookie := http.Cookie{
				Name:     "user_id",
				Value:    strconv.Itoa(user.ID),
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
			}
			SetUpConnect()
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
			return
		}
	}
	StockMessageErreurConnexion = "Identifiants incorrects"

	http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
}
