package main

import (
	"github.com/michael-golfi/Grott/grott"
	"github.com/michael-golfi/Grott/alfredo/controller"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/dialog"
	"log"
)

func main() {

	s := &dialog.SimpleDialog{}
	dialogs := []types.Dialoger{s}
	router := dialog.NewInMemoryStorageRouter(dialogs)
	c := &controller.Controller{
		Router: router,
	}
	log.Fatal(grott.ListenAndServe(c, router))
}