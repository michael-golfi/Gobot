package controller

import (
	"github.com/michael-golfi/Grott/grott/types"
)

var (
	send = make(chan types.Message, 1)
	receive = make(chan types.Message, 1)
)

func StartBot(d types.Dialoger) {
	types.Start(d, send, receive)
}