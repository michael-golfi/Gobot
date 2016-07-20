package controller

import (
	"errors"
	"fmt"
	"github.com/michael-golfi/Grott/grott/dialog"
	"github.com/michael-golfi/Grott/grott/types"
	"log"
)

type Controller struct {
	Router *dialog.DialogRouter
}

func (c Controller) Post(msg *types.Activity) {
	if msg.Type == "message" {
		c.Router.HandleMessage(msg)
	}

	c.HandleSystemMessage(msg)
}

func (c Controller) HandleSystemMessage(msg *types.Activity) {
	fmt.Println("HandleSystemMessage")
	fmt.Println(msg.Type)

	var out string
	switch msg.Type {
	case "BotAddedToConversation":
		out = "BotAddedToConversation"
	case "UserAddedToConversation":
		out = "UserAddedToConversation"
	case "BotRemovedFromConversation":
		out = "BotRemovedFromConversation"
	case "UserRemovedFromConversation":
		out = "UserRemovedFromConversation"
	case "EndOfConversation":
		out = "EndOfConversation"
	default:
		out = "Not A System Message"
	}
	log.Printf("HandleSystemMessage: %s", out)
}
