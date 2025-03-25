package Fonctions

import (
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
}
