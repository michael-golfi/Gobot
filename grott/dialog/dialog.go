package dialog

import (
	"github.com/michael-golfi/Grott/grott/activity"
	"github.com/michael-golfi/Grott/grott/storage"
)

type Dialoger interface {
	MessageReceived(ctx *storage.DialogContext, msg *activity.Activity)
}


