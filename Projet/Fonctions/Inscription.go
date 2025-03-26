package Fonctions

import (
	"fmt"
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

var StockMessageErreurInscription string

type DataInscrition struct {
	MessageErreur string
	IsConnect     bool
	User          db.DB
}

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		InscriptionGet(w, r)

	case "POST":
		InscriptionPost(w, r)
	}
}

func InscriptionGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/inscription.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
		return
	}

	data := DataInscrition{
		MessageErreur: StockMessageErreurInscription,
		IsConnect:     IsConnectStock,
		User:          db.GetUser(cookie.Value),
	}

	StockMessageErreurInscription = ""

	tmpl.Execute(w, data)
}

func InscriptionPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // Gère s'il y a des erreurs
		fmt.Printf("Erreur ParseForm() : %v\n", err)
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	prenom := r.Form.Get("prenom")
	nom := r.Form.Get("nom")
	email := r.Form.Get("email")
	mdp := r.Form.Get("mdp")

	//fmt.Println(prenom, nom, email, mdp)

	tabUser := db.GetDB()

	for _, user := range tabUser {
		if user.Email == email {
			StockMessageErreurInscription = "Adresse mail déjà utilisé"
			http.Redirect(w, r, "/Inscription", http.StatusSeeOther)
			return
		}
	}

	db.InsertUser(prenom, nom, mdp, email)

	http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
}
