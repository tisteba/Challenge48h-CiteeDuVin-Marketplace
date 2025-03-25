package Fonctions

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Day struct {
	Date           time.Time
	IsCurrent      bool
	IsToday        bool
	HasAppointment bool
}

type MonthData struct {
	Year      int
	Month     time.Month
	MonthName string
	Days      []Day
	PrevMonth int
	PrevYear  int
	NextMonth int
	NextYear  int
	TimeSlots []string
	Now       time.Time
}

func RdvPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		calendarHandler(w, r)
	case "POST":
		bookHandler(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	year, month, _ := now.Date()

	// Gestion des paramètres d'URL
	queryYear := r.URL.Query().Get("year")
	queryMonth := r.URL.Query().Get("month")

	if queryYear != "" && queryMonth != "" {
		var err error
		year, err = strconv.Atoi(queryYear)
		if err != nil {
			http.Error(w, "Année invalide", http.StatusBadRequest)
			return
		}
		monthInt, err := strconv.Atoi(queryMonth)
		if err != nil || monthInt < 1 || monthInt > 12 {
			http.Error(w, "Mois invalide", http.StatusBadRequest)
			return
		}
		month = time.Month(monthInt)
	}

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	// Calcul des dates à afficher
	startDate := firstOfMonth.AddDate(0, 0, -int(firstOfMonth.Weekday()))
	if firstOfMonth.Weekday() == time.Sunday {
		startDate = firstOfMonth
	}
	endDate := lastOfMonth.AddDate(0, 0, 6-int(lastOfMonth.Weekday()))

	// Préparation des données pour le template
	data := MonthData{
		Year:      year,
		Month:     month,
		MonthName: month.String(),
		PrevMonth: int(month) - 1,
		PrevYear:  year,
		NextMonth: int(month) + 1,
		NextYear:  year,
		TimeSlots: generateTimeSlots(),
		Now:       now,
	}

	// Ajustement pour janvier et décembre
	if month == time.January {
		data.PrevMonth = 12
		data.PrevYear = year - 1
	} else if month == time.December {
		data.NextMonth = 1
		data.NextYear = year + 1
	}

	// Remplissage des jours du calendrier
	currentDate := startDate
	for !currentDate.After(endDate) {
		day := Day{
			Date:      currentDate,
			IsCurrent: currentDate.Month() == month,
			IsToday:   currentDate.Year() == now.Year() && currentDate.Month() == now.Month() && currentDate.Day() == now.Day(),
		}
		data.Days = append(data.Days, day)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("../Front/HTML/rdv.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur de chargement du template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Erreur d'exécution du template: %v", err), http.StatusInternalServerError)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	// Ici vous pourriez traiter les données du formulaire
	// date := r.FormValue("date")
	// time := r.FormValue("time")
	// etc.

	// Redirection après traitement
	http.Redirect(w, r, "/PriseRdv", http.StatusSeeOther)
}

func generateTimeSlots() []string {
	slots := make([]string, 0)
	start, _ := time.Parse("15:04", "09:00")
	end, _ := time.Parse("15:04", "18:00")

	for current := start; current.Before(end); current = current.Add(30 * time.Minute) {
		slots = append(slots, current.Format("15:04"))
	}

	return slots
}
