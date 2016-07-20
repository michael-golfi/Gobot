package types

type Dialoger interface {
	MessageReceived(ctx *DialogContext, msg *Activity) error
	CalculateScore(msg *Activity) (int, error)
}

type DialogContext struct {
	ConversationData          map[string]string
	PerUserInConversationData map[string]string
	UserData                  map[string]string
}
