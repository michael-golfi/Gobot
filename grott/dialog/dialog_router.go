package dialog

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/storage"
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/storage/inmemory"
)

type DialogRouter struct {
	Dialogs        []types.Dialoger
	ContextStorage storage.ContextStorage
}

func NewInMemoryStorageRouter(dialogs []types.Dialoger) *DialogRouter {
	inMemoryStorage := inmemory.NewInMemoryStorage()
	return NewRouter(dialogs, inMemoryStorage)
}

func NewRouter(dialogs []types.Dialoger, storage storage.ContextStorage) *DialogRouter {
	return &DialogRouter{
		Dialogs:        dialogs,
		ContextStorage: storage,
	}
}

func (router *DialogRouter) HandleMessage(message *types.Activity) {

	go func(d *DialogRouter, m *types.Activity) {

		msgCtx, err := d.ContextStorage.Get(m.Conversation.Id)

		if err != nil {
			fmt.Println(err.Error())

			d.ContextStorage.Save(m.Conversation.Id, &types.DialogContext{
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