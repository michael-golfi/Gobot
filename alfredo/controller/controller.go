package controller

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/dialog"
	"log"
	"github.com/michael-golfi/Grott/grott/activity"
)

type Controller struct {
	Router *dialog.DialogRouter
}

func (c Controller) Post(msg *activity.Activity) {
	if msg.Type == "message" {
		c.Router.HandleMessage(msg)
	} else {
		c.HandleSystemMessage(msg)
	}
}

func (c Controller) HandleSystemMessage(msg *activity.Activity) {
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
