package Fonctions

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Wine représente la structure des données du fichier JSON
type Wine struct {
	Points              int     `json:"points"`
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	TasterName          *string `json:"taster_name"`
	TasterTwitterHandle *string `json:"taster_twitter_handle"`
	Price               int     `json:"price"`
	Designation         *string `json:"designation"`
	Variety             string  `json:"variety"`
	Region1             *string `json:"region_1"`
	Region2             *string `json:"region_2"`
	Province            string  `json:"province"`
	Country             string  `json:"country"`
	Winery              string  `json:"winery"`
}

// LoadWines lit le fichier JSON et le parse
func LoadWines() ([]Wine, error) {
	file, err := os.Open("../Data/wine-data-set.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var wines []Wine
	if err := json.Unmarshal(bytes, &wines); err != nil {
		return nil, err
	}

	return wines, nil
}
