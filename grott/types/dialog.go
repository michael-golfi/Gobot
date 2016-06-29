package types

import (
	"fmt"
)

type Dialoger interface {
	SetInput(in chan Message)
	SetOutput(out chan Message)

	GetInput() chan Message
	GetOutput() chan Message

	//Begin(session SessionKeeper, args interface{})
	//ReplyReceived(session SessionKeeper)
	//DialogResumed(session SessionKeeper, result chan interface{})
	//CompareConfidence(action SessionAction, language, utterance string, score int)
}

type DialogAction interface {

}

type DialogContext struct {

}

func Start(dialog Dialoger, in, out chan Message) {
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