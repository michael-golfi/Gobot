package dialog

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
)

type SimpleDialog struct {
	In  chan types.Message
	Out chan types.Message
}

func NewSimpleDialog(input, output chan types.Message) *SimpleDialog {
	simpleDialog := &SimpleDialog{
		In:  input,
		Out: output,
	}

	go simpleDialog.startListening()

	return simpleDialog
}

func (d *SimpleDialog) startListening() {
	for {
		message := <-d.In

		// TODO - Get Session

		fmt.Println(message.Text)
	}
}

func (d *SimpleDialog) messageReceived(session *types.Session) {

}
