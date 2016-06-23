package controller

import (
	"github.com/michael-golfi/Grott/grott/types"
	"fmt"
	"github.com/michael-golfi/Grott/grott/dialog"
	"errors"
)

var (
	send chan types.Message
	receive chan types.Message
)

func init() {
	send = make(chan types.Message, 5)
	receive = make(chan types.Message, 5)
	dialog.NewSimpleDialog(send, receive)
}

func Post(msg *types.Message) *types.Message {
	if msg.MessageType == "Message" {
		send <- *msg
		msgReceived := <-receive
		return &msgReceived
	}

	m, err := handleSystemMessage(msg)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	return m
}

func handleSystemMessage(msg *types.Message) (*types.Message, error) {
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