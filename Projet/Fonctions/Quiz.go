package Fonctions

import (
	"html/template"
	"net/http"
)

type DataQuiz struct {
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

	data := DataQuiz{}
	tmpl.Execute(w, data)
}

func QuizPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Quiz", http.StatusSeeOther)
}
