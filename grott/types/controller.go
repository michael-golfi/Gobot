package types

type Controller interface {
	Post(msg *Message) (*Message, error)
	HandleSystemMessage(msg *Message) (*Message, error)
}