package grott

import "time"

type Session struct {
	Options      SessionOptions
	MessageSent  bool
	isReset      bool
	LastSentTime time.Time
	SendQueue    chan Message
	Dialogs      []*Dialog
}

type SessionAction struct {

}