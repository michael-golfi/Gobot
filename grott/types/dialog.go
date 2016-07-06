package types

import (
	"fmt"
)

type Dialoger interface {
	MessageReceived(ctx DialogContext, msg Message) (Message, error)
	CalculateScore(msg Message) (int, error)
}

type DialogAction interface {

}

type DialogContext struct {
	ConversationData          map[string]interface{}
	PerUserInConversationData map[string]interface{}
	UserData                  map[string]interface{}
}

func Start(dialog Dialoger, in, out chan Message) {
	ctx := DialogContext{
		ConversationData: make(map[string]interface{}),
		UserData: make(map[string]interface{}),
		PerUserInConversationData: make(map[string]interface{}),
	}

	for {
		message := <-in

		// TODO - Get Session
		// Message received

		msg, err := dialog.MessageReceived(ctx, message)

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		out <- msg

		fmt.Println(message.Text)
	}
}