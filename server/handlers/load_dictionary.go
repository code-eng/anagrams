package handlers

import (
	"anagrams/server/state"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadDict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Load dict called with %v method", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Load dict called with %v method", r.Method)
		http.Error(w, "Can't read body", http.StatusInternalServerError)
		return
	}

	var parsedWords []string
	err = json.Unmarshal(body, &parsedWords)
	if err != nil {
		http.Error(w, "Body should be array of strings", http.StatusBadRequest)
		return
	}

	state.LoadDictionary(parsedWords)

	log.Printf("Loaded %v words\n", len(parsedWords))
}
