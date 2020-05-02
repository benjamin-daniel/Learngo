package reader

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AnimalFacts struct {
	CommonName     string   `json:"commonName"`
	ScientificName string   `json:"scientificName"`
	HeightInInches int      `json:"heightInInches"`
	FavoriteFoods  []string `json:"favoriteFoods"`
	CanSwim        bool     `json:"canSwim"`
}

// SlothFacts is used to get
func SlothFacts() {
	jsonData := `
{
    "commonName":     "brown-throated three-toed sloth",
    "scientificName": "Bradypus variegatus",
    "heightInInches": 31,
    "favoriteFoods":  ["Cecropia leaves", "Hibiscus flowers"],
    "canSwim":        true
}`

	// an io.Reader that isn't a file!
	rdr := strings.NewReader(jsonData)

	var sloth AnimalFacts
	if err := json.NewDecoder(rdr).Decode(&sloth); err != nil {
		fmt.Printf("error deserializing JSON: %v", err)
		return
	}
	fmt.Println(&sloth)

	fmt.Printf("Hi! I'm a %s (%s)! I'm about %d\" tall, and love eating %v!\n",
		sloth.CommonName, sloth.ScientificName, sloth.HeightInInches,
		sloth.FavoriteFoods)
	if sloth.CanSwim {
		fmt.Println("By the way, I can swim!")
	}
}
