package Fonctions

import (
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

type DataFrance struct {
	WineList  []Wine
	IsConnect bool
	User      db.DB
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

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}

	data := DataFrance{
		WineList:  RecupWineContry("France"),
		IsConnect: IsConnectStock,
		User:      db.GetUser(cookie.Value),
	}
	tmpl.Execute(w, data)
}

func VinsFrancePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/VinsDeFrance", http.StatusSeeOther)
}
