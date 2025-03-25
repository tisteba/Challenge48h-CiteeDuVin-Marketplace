package Fonctions

import (
	"html/template"
	"net/http"
)

type DataFrance struct {
	WineList []Wine
}

func VinsFrancePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		VinsFranceGet(w, r)

	case "POST":
		VinsFrancePost(w, r)
	}
}

func VinsFranceGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/vinsfrance.html"))

	data := DataFrance{
		WineList: RecupWineContry("France"),
	}
	tmpl.Execute(w, data)
}

func VinsFrancePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/VinsDeFrance", http.StatusSeeOther)
}
