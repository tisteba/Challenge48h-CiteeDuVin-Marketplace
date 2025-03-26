package Fonctions

import (
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

type VinsPays struct {
	Pays     string
	WineList []Wine // Changé pour être une slice de Wine
}

type DataMonde struct {
	Vins      []VinsPays
	IsConnect bool
	User      db.DB
}

func VinsMondePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		VinsMondePageGet(w, r)
	case "POST":
		VinsMondePagePost(w, r)
	}
}

func VinsMondePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/vinsmonde.html"))
	var Final []VinsPays

	var pays = [5]string{"France", "US", "Portugal", "Italy", "Spain"}

	for i := 0; i < len(pays); i++ {
		var Instance VinsPays
		Instance.Pays = pays[i]

		wines := RecupWineContry(pays[i])

		// Limiter à 3 vins maximum
		if len(wines) > 3 {
			Instance.WineList = wines[:3] // Prendre seulement les 3 premiers vins
		} else {
			Instance.WineList = wines // Prendre tous les vins s'il y en a moins de 3
		}

		Final = append(Final, Instance)
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
		return
	}

	data := DataMonde{
		Vins:      Final,
		IsConnect: IsConnectStock,
		User:      db.GetUser(cookie.Value),
	}
	tmpl.Execute(w, data)
}

func VinsMondePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/VinsDuMonde", http.StatusSeeOther)
}
