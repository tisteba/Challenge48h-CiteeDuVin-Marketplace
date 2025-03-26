package Fonctions

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Handler pour la page Box
func BoxPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("../Front/HTML/box.html"))
		tmpl.Execute(w, nil)
	}
}

// Handler pour l'API de sélection aléatoire
func GetLootBox(w http.ResponseWriter, r *http.Request) {
	// 1. Parser les paramètres
	budgetStr := r.URL.Query().Get("budget")
	countStr := r.URL.Query().Get("count")
	budget, _ := strconv.Atoi(budgetStr)
	count, _ := strconv.Atoi(countStr)

	// 2. Charger et filtrer les vins
	wines, err := LoadWines()
	if err != nil {
		http.Error(w, "Erreur de chargement des vins", http.StatusInternalServerError)
		return
	}

	filtered := filterWinesForBox(wines, budget, count)

	// 3. Renvoyer la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

// Logique de sélection
func filterWinesForBox(wines []Wine, maxTotalPrice int, count int) []Wine {
	rand.Seed(time.Now().UnixNano())
	var selected []Wine
	remainingBudget := maxTotalPrice

	// Filtre initial : vins avec score > 85 ET prix > 0 ET prix <= budget
	var candidates []Wine
	for _, wine := range wines {
		if wine.Points >= 85 && wine.Price > 0 && wine.Price <= remainingBudget {
			candidates = append(candidates, wine)
		}
	}

	// ... (le reste de la fonction reste identique)

	// Sélection aléatoire
	for len(selected) < count && len(candidates) > 0 {
		index := rand.Intn(len(candidates))
		selectedWine := candidates[index]

		if remainingBudget >= selectedWine.Price {
			selected = append(selected, selectedWine)
			remainingBudget -= selectedWine.Price
			candidates = append(candidates[:index], candidates[index+1:]...)
		}
	}
	fmt.Println(selected)

	return selected
}
