package main

import (
	"log"
	"github.com/michael-golfi/Grott/grott"
)

func main() {
	log.Print("Starting Bot")
	log.Fatal(grott.ListenAndServe())
}