package main

import (
	"github.com/michael-golfi/Grott/grott"
	"log"
)

func main() {
	log.Print("Starting Bot")
	log.Fatal(grott.ListenAndServe())
}
