package dialog

import (
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/storage"
	"sync"
	"fmt"
)

type DialogRouter struct {
	Dialogs        []types.Dialoger
	ContextStorage storage.ContextStorage
}

func NewInMemoryStorageRouter(dialogs []types.Dialoger) *DialogRouter {
	inMemoryStorage := storage.NewInMemoryStorage()
	return NewRouter(dialogs, inMemoryStorage)
}

func NewRouter(dialogs []types.Dialoger, storage storage.ContextStorage) *DialogRouter {
	return &DialogRouter{
		Dialogs: dialogs,
		ContextStorage: storage,
	}
}

func (router *DialogRouter) HandleMessage(message *types.Activity) error {

	var e error

	wait := sync.WaitGroup{}
	wait.Add(1)

	go func(d *DialogRouter, m *types.Activity) {
		highestIndex := 0
		var err error

		if len(d.Dialogs) > 1 {
			highestIndex, err = getHighestScoringDialog(d.Dialogs, m)

			if err != nil {
				e = err
				wait.Done()
				return
			}
		}

		msgCtx, err := d.ContextStorage.Get(m.Conversation.Id)
		if err != nil {
			fmt.Println(err.Error())

			d.ContextStorage.Save(m.Conversation.Id, &types.DialogContext{
				ConversationData:          make(map[string]string, 1),
				PerUserInConversationData: make(map[string]string, 1),
				UserData:                  make(map[string]string, 1),
			})
			// TODO - Just means the ctx doesn't exist
		}

		err = d.Dialogs[highestIndex].MessageReceived(msgCtx, m)
		if err != nil {
			e = err
			wait.Done()
			return
		}

		e = nil
		wait.Done()
	}(router, message)

	wait.Wait()
	return e
}

func getHighestScoringDialog(dialogs []types.Dialoger, msg *types.Activity) (int, error) {

	highestIndex := 0
	highestScore := 0

	for i, dialog := range dialogs {

		score, err := dialog.CalculateScore(msg)

		if err != nil {
			return -1, err
		}

		if score > highestScore {
			highestScore = score
			highestIndex = i
		}

	}

	return highestIndex, nil
}