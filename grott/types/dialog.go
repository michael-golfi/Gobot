package types

import (
	"fmt"
)

type Dialog interface {
	SetInput(in chan Message)
	SetOutput(out chan Message)

	GetInput() chan Message
	GetOutput() chan Message

	Begin(session Session, args interface{})
	ReplyReceived(session Session)
	DialogResumed(session Session, result chan interface{})
	CompareConfidence(action SessionAction, language, utterance string, score int)
}

type DialogAction interface {
}

func Start(dialog Dialog, in, out chan Message) {
	dialog.SetInput(in)
	dialog.SetOutput(out)

	go func(In, Out chan Message) {
		for {
			message := <-In

			// TODO - Get Session

			fmt.Println(message.Text)
		}
	}(dialog.GetInput(), dialog.GetOutput())
}