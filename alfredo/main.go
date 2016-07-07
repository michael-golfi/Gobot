package main

import (
	"github.com/michael-golfi/Grott/grott"
	"github.com/michael-golfi/Grott/alfredo/controller"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/dialog"
)

func main() {


	c := &controller.Controller{}
	s := &dialog.SimpleDialog{}
	dialogs := []types.Dialoger{s}
	router := dialog.NewInMemoryStorageRouter(dialogs)
	grott.ListenAndServe(c, router)
}