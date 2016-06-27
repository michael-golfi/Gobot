package controller

import (
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/dialog"
)

var (
	send = make(chan types.Message, 1)
	receive = make(chan types.Message, 1)
)

func StartBot(d *types.Dialog) {
	dialog.Start(d, send, receive)
}