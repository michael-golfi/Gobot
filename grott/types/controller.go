package types

type Controller interface {
	Post(msg *Message) *Message
	HandleSystemMessage(msg *Message) (*Message, error)
}