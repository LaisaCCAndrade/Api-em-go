package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Species  string `json:"species"`
	Type     string `json:"type"`
	Gender   string `json:"gender"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
}

type CharactersResponse struct {
	Results []Character `json:"results"`
}

func getCharacters(w http.ResponseWriter, r *http.Request) {
	url := "https://rickandmortyapi.com/api/character"
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var charactersResponse CharactersResponse
	if err := json.NewDecoder(response.Body).Decode(&charactersResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(charactersResponse.Results)
}

func main() {
	http.HandleFunc("/characters", getCharacters)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
