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

func (c Controller) Post(msg *types.Message) (*types.Message, error) {
	if msg.MessageType == "Message" {
		return c.Router.HandleMessage(msg)
	}

	return c.HandleSystemMessage(msg)
}

func (c Controller) HandleSystemMessage(msg *types.Message) (*types.Message, error) {
	fmt.Println("HandleSystemMessage")
	fmt.Println(msg.MessageType)
	if msg.MessageType == "BotAddedToConversation" {
		fmt.Println("BotAddedToConversation")
		return &types.Message{
			Text: "Hello World",
		}, nil
	}

	if msg.MessageType == "UserAddedToConversation" {
		fmt.Println("UserAddedToConversation")
		return &types.Message{
			Text: "Hello Human!",
		}, nil
	}

	if msg.MessageType == "BotRemovedFromConversation" {
		fmt.Println("BotRemovedFromConversation")
		return &types.Message{
			Text: "Goodbye Bot!",
		}, nil
	}

	if msg.MessageType == "UserRemovedFromConversation" {
		fmt.Println("UserRemovedFromConversation")
		return &types.Message{
			Text: "Goodbye Human!",
		}, nil
	}

	if msg.MessageType == "EndOfConversation" {
		fmt.Println("EndOfConversation")
		return &types.Message{
			Text: "End of Conversation",
		}, nil
	}

	return nil, errors.New("Not a System Message")
}

