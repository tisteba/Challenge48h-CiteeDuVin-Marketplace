package Fonctions

import (
	"html/template"
	"net/http"
)

type VinsPays struct {
	Pays     string
	WineList []Wine // Changé pour être une slice de Wine
}

type DataMonde struct {
	IsConnect bool
	Vins      []VinsPays
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

	data := DataMonde{
		Vins:      Final,
		IsConnect: false,
	}
	tmpl.Execute(w, data)
}

func VinsMondePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/VinsDuMonde", http.StatusSeeOther)
}
