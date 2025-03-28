package Fonctions

import (
	"fmt"
	db "marketplace/DataBase"
	"net/http"
)

func SetHttp() {
	fs := http.FileServer(http.Dir("../Front/CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", RedirectBasePage)
	http.HandleFunc("/HomePage", HomePage)
	http.HandleFunc("/TousLesVins", EnsembleVinsPage)
	http.HandleFunc("/VinsDeFrance", VinsFrancePage)
	http.HandleFunc("/VinsDuMonde", VinsMondePage)
	http.HandleFunc("/NosPepites", PepitesPage)
	http.HandleFunc("/Quiz", QuizPage)
	http.HandleFunc("/Connexion", ConnexionPage)
	http.HandleFunc("/Inscription", InscriptionPage)
	http.HandleFunc("/Encheres", EncheresPage)
	http.HandleFunc("/PriseRdv", RdvPage)
	http.HandleFunc("/Box", BoxPage)
	http.HandleFunc("/api/lootbox", GetLootBox)
	http.HandleFunc("/Panier", PanierPage)

	SetDB()
}

func SetDB() {
	db.CreateTableUsers()
	fmt.Println("")
}
