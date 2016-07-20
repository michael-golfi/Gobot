package main

import (
	"github.com/michael-golfi/Grott/alfredo/controller"
	"github.com/michael-golfi/Grott/alfredo/dialog"
	"github.com/michael-golfi/Grott/grott"
	"github.com/michael-golfi/Grott/grott/dialog"
	"github.com/michael-golfi/Grott/grott/types"
	"log"
)

func main() {
	s := &bot_dialog.SimpleDialog{}
	dialogs := []types.Dialoger{s}
	router := dialog.NewInMemoryStorageRouter(dialogs)
	c := &controller.Controller{
		Router: router,
	}
	log.Fatal(grott.ListenAndServe(c, router))
}
