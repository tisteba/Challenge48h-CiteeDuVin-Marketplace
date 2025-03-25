package Fonctions

import (
	"net/http"
)

func SetHttp() {
	fs := http.FileServer(http.Dir("../Front/CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", RedirectBasePage)
	http.HandleFunc("/HomePage", HomePage)
}
