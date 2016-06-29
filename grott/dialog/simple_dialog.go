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
*
func (d SimpleDialog) Begin(sessionKeeper types.SessionKeeper, args interface{}) {

}

func (d SimpleDialog) ReplyReceived(sessionKeeper types.SessionKeeper) {

}

func (d SimpleDialog) DialogResumed(sessionKeeper types.SessionKeeper, result chan interface{}) {

}*/