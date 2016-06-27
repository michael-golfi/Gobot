package main

import (
	"github.com/michael-golfi/Grott/grott"
	"log"
	"github.com/michael-golfi/Grott/grott/dialog"
)

func main() {
	log.Print("Starting Bot")

	dialog := new (dialog.SimpleDialog)

	log.Fatal(grott.ListenAndServe(*dialog))
}