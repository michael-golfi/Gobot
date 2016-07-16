package types

type Controller interface {
	Post(msg *Activity) error
	HandleSystemMessage(msg *Activity) error
}