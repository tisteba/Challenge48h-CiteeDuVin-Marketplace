package Fonctions

import (
	"html/template"
	db "marketplace/DataBase"
	"net/http"
)

type DataQuiz struct {
	IsConnect bool
	User      db.DB
}

func QuizPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		QuizGet(w, r)

	case "POST":
		QuizPost(w, r)
	}
}

func QuizGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../Front/HTML/quiz.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Connexion", http.StatusSeeOther)
		return
	}

	data := DataQuiz{
		IsConnect: IsConnectStock,
		User:      db.GetUser(cookie.Value),
	}
	tmpl.Execute(w, data)
}

func QuizPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Quiz", http.StatusSeeOther)
}
