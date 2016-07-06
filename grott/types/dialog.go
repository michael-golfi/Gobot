package types

import (
	"fmt"
)

type Dialoger interface {
	MessageReceived(ctx *DialogContext, msg *Message) (*Message, error)
	CalculateScore(msg *Message) (int, error)
}

type DialogContext struct {
	ConversationData          map[string]string
	PerUserInConversationData map[string]string
	UserData                  map[string]string
}

func Start(dialog Dialoger, in, out chan Message) {
	ctx := &DialogContext{
		ConversationData: make(map[string]string),
		UserData: make(map[string]string),
		PerUserInConversationData: make(map[string]string),
	}

	for {
		message := <-in

		// TODO - Get Session
		// Message received

		msg, err := dialog.MessageReceived(ctx, &message)

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		out <- *msg

		fmt.Println(message.Text)
	}
}