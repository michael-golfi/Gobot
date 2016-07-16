package controller

import (
	"fmt"
	"errors"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/dialog"
)

type Controller struct {
	Router *dialog.DialogRouter
}

func (c Controller) Post(msg *types.Activity) error {
	if msg.Type == "message" {
		return c.Router.HandleMessage(msg)
	}

	return c.HandleSystemMessage(msg)
}

func (c Controller) HandleSystemMessage(msg *types.Activity) error {
	fmt.Println("HandleSystemMessage")
	fmt.Println(msg.Type)
	if msg.Type == "BotAddedToConversation" {
		fmt.Println("BotAddedToConversation")
		return nil
	}

	if msg.Type == "UserAddedToConversation" {
		fmt.Println("UserAddedToConversation")
		return nil
	}

	if msg.Type == "BotRemovedFromConversation" {
		fmt.Println("BotRemovedFromConversation")
		return nil
	}

	if msg.Type == "UserRemovedFromConversation" {
		fmt.Println("UserRemovedFromConversation")
		return nil
	}

	if msg.Type == "EndOfConversation" {
		fmt.Println("EndOfConversation")
		return nil
	}

	return errors.New("Not a System Message")
}

