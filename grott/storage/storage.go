package storage

type DialogContext struct {
	ConversationData          map[string]string
	PerUserInConversationData map[string]string
	UserData                  map[string]string
}

type ContextStorage interface {
	Get(id string) (*DialogContext, error)
	Save(id string, data *DialogContext) error
	Delete(id string)
}
