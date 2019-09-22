package main

import (
	"log"
)
import "anagrams/server"

func main() {
	log.Fatal(server.NewServer().ListenAndServe())
}
