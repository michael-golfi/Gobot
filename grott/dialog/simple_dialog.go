package dialog

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
)

type SimpleDialog struct {
	In  chan types.Message
	Out chan types.Message
}

func Start(dialog *SimpleDialog, in, out chan types.Message) {
	(*dialog).SetInput(in)
	(*dialog).SetOutput(out)

	go func(In, Out chan types.Message) {
		for {
			message := <-In

			// TODO - Get Session

			fmt.Println(message.Text)
		}
	}(dialog.In, dialog.Out)
}

func (d *SimpleDialog) SetInput(input chan types.Message) {
	d.In = input
}

func (d *SimpleDialog) SetOutput(output chan types.Message) {
	d.Out = output
}

/**
	Interface functions
 */
func (d *SimpleDialog) Begin(session types.Session, args interface{}) {

}

func (d *SimpleDialog) ReplyReceived(session types.Session) {

}

func (d *SimpleDialog) DialogResumed(session types.Session, result chan interface{}) {

}

func (d *SimpleDialog) CompareConfidence(action types.SessionAction, language, utterance string, score int) {

}
