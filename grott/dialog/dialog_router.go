package dialog

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/storage"
	"github.com/michael-golfi/Grott/grott/activity"
	"github.com/michael-golfi/Grott/grott/storage/inmemory"
)

type DialogRouter struct {
	Dialogs        []Dialoger
	ContextStorage storage.ContextStorage
}

func NewInMemoryStorageRouter(dialogs []Dialoger) *DialogRouter {
	inMemoryStorage := inmemory.NewInMemoryStorage()
	return NewRouter(dialogs, inMemoryStorage)
}

func NewRouter(dialogs []Dialoger, storage storage.ContextStorage) *DialogRouter {
	return &DialogRouter{
		Dialogs:        dialogs,
		ContextStorage: storage,
	}
}

func (router *DialogRouter) HandleMessage(message *activity.Activity) {

	go func(d *DialogRouter, m *activity.Activity) {

		msgCtx, err := d.ContextStorage.Get(m.Conversation.Id)

		if err != nil {
			fmt.Println(err.Error())

			d.ContextStorage.Save(m.Conversation.Id, &storage.DialogContext{
				ConversationData:          make(map[string]string, 1),
				PerUserInConversationData: make(map[string]string, 1),
				UserData:                  make(map[string]string, 1),
			})
		}

		for _, dialog := range d.Dialogs {
			// TODO - Context should be checked for thread safety
			go dialog.MessageReceived(msgCtx, m)
		}

	}(router, message)

}