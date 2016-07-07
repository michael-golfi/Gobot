package controller

import (
	"fmt"
	"errors"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/dialog"
)

type Controller struct {
	router dialog.DialogRouter
}

func (c Controller) Post(msg *types.Message) *types.Message {
	if msg.MessageType == "Message" {
		msg, err := c.router.HandleMessage(msg)

		// TODO - Bubble up error
		if err != nil {
			fmt.Errorf("Post Error: %s", err.Error())
		}

		return msg
	}

	m, err := c.HandleSystemMessage(msg)

	// TODO - Bubble up error
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	return m
}

func (c Controller) HandleSystemMessage(msg *types.Message) (*types.Message, error) {
	if msg.MessageType == "BotAddedToConversation" {
		return &types.Message{
			Text: "Hello World",
		}, nil
	}

	if msg.MessageType == "UserAddedToConversation" {
		return &types.Message{
			Text: "Hello Human!",
		}, nil
	}

	if msg.MessageType == "BotRemovedFromConversation" {
		return &types.Message{
			Text: "Goodbye Bot!",
		}, nil
	}

	if msg.MessageType == "UserRemovedFromConversation" {
		return &types.Message{
			Text: "Goodbye Human!",
		}, nil
	}

	if msg.MessageType == "EndOfConversation" {
		return &types.Message{
			Text: "End of Conversation",
		}, nil
	}
	return nil, errors.New("Not a System Message")
}

