package types

type Dialoger interface {
	MessageReceived(ctx *DialogContext, msg *Activity)
}

type DialogContext struct {
	ConversationData          map[string]string
	PerUserInConversationData map[string]string
	UserData                  map[string]string
}
