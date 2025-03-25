package main

import (
	"fmt"
	fc "marketplace/Fonctions"
	"net/http"
)

func main() {
	fc.SetHttp()

	wines, err := fc.LoadWines()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier JSON :", err)
		return
	}

	for _, wine := range wines {
		if wine.Country == "US" {
			fmt.Printf("Titre: %s | Points: %d | Prix: %d$\n", wine.Title, wine.Points, wine.Price)
		}
	}

	fmt.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
	fmt.Println("")
}
