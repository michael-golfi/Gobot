package dialog

import (
	"github.com/michael-golfi/Grott/grott/types"
)

type SimpleDialog struct {
	in  chan types.Message
	out chan types.Message
}

func (d SimpleDialog) SetInput(input chan types.Message) {
	d.in = input
}

func (d SimpleDialog) SetOutput(output chan types.Message) {
	d.out = output
}

func (d SimpleDialog) GetInput() chan types.Message {
	return d.in
}

func (d SimpleDialog) GetOutput() chan types.Message {
	return d.out
}

/**
	Interface functions
 */
func (d SimpleDialog) Begin(session types.Session, args interface{}) {

}

func (d SimpleDialog) ReplyReceived(session types.Session) {

}

func (d SimpleDialog) DialogResumed(session types.Session, result chan interface{}) {

}

func (d SimpleDialog) CompareConfidence(action types.SessionAction, language, utterance string, score int) {

}