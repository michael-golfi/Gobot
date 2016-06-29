package dialog

import (
	"github.com/michael-golfi/Grott/grott/types"
)

type SimpleDialog struct {}

func (d SimpleDialog) MessageReceived(ctx types.DialogContext, msg types.Message) (types.Message, error) {
	return types.Message{
		Text: "Hello World!",
	}
}