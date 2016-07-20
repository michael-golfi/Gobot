package types

type Controller interface {
	Post(msg *Activity)
	HandleSystemMessage(msg *Activity)
}
