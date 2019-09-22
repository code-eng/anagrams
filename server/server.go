package server

import (
	"anagrams/server/handlers"
	"net/http"
	"time"
)

func addRoutes() {
	http.HandleFunc("/load", handlers.LoadDict)
	http.HandleFunc("/get", handlers.GetAnagrams)
}

func NewServer() *http.Server {
	addRoutes()
	return &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
