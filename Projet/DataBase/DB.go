package DB

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type DB struct {
	ID     int
	Prenom string
	Nom    string
	Email  string
	Mdp    string
	Solde  int
}

func connect_DB() *sql.DB {
	dbUsers, err := sql.Open("sqlite", "../DataBase/Stockage/DB.db")
	if err != nil {
		panic(err)
	}

	// Vérifier la connexion
	if err := dbUsers.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données DB !")
	}
	return dbUsers
}

func CreateTableUsers() {
	db := connect_DB()
	defer db.Close()

	// Structure de notre base de donnée
	query := ` 
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		prenom TEXT NOT NULL,
		nom TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		solde IINTEGER DEFAULT 0
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table 'user' créée avec succès.")
}

func GetDB() []DB {
	db := connect_DB()
	defer db.Close()

	var donnes []DB
	rows, err := db.Query("SELECT id, prenom, nom, email, password, solde FROM user;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var base DB
		if err := rows.Scan(&base.ID, &base.Prenom, &base.Nom, &base.Email, &base.Mdp, &base.Solde); err != nil {
			panic(err)
		}

		donnes = append(donnes, base)
	}
	return donnes
}

func InsertUser(prenom string, nom string, password string, email string) {
	db := connect_DB()

	query := `INSERT INTO users (prenom, nom, email, mdp) VALUES (?, ?, ?, ?);`
	_, err := db.Exec(query, prenom, nom, email, password)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
}

func ChangeSolde(Ajout string, userID int) {
	db := connect_DB()
	defer db.Close()

	query := "UPDATE users SET solde = ? WHERE id = ?"
	_, err := db.Exec(query, Ajout, userID)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du mot de passe :", err)
	}
}
