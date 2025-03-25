package Fonctions

import (
	"net/http"
)

func SetHttp() {
	fs := http.FileServer(http.Dir("./CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/HomePage", HomePage)
}
