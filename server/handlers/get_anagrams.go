package handlers

import (
	"anagrams/server/state"
	"encoding/json"
	"log"
	"net/http"
)

func GetAnagrams(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()

	word := query.Get("word")
	if word == "" {
		http.Error(w, "word is not provided", http.StatusBadRequest)
		return
	}

	anagrams := state.GetAnagrams(word)
	if len(anagrams) == 0 {
		anagrams = nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(anagrams)
	if err != nil {
		http.Error(w, "Cannot encode anagrams", http.StatusInternalServerError)
		return
	}

	log.Printf("Got anagrmas for %v ", word)
}
