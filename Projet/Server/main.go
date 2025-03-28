package main

import (
	"fmt"
	fc "marketplace/Fonctions"
	"net/http"
)

func main() {
	fc.SetHttp()

	fmt.Println("\nServeur démarré sur http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
	fmt.Println("")
}
