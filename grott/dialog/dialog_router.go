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
	inMemoryStorage := storage.InMemoryStorage{
		Cache: make(map[string]*types.DialogContext, 1),
	}
	return NewRouter(dialogs, inMemoryStorage)
}

func NewRouter(dialogs []types.Dialoger, storage storage.ContextStorage) *DialogRouter {
	return &DialogRouter{
		Dialogs: dialogs,
		ContextStorage: storage,
	}
}

func (router *DialogRouter) HandleMessage(message *types.Message) (*types.Message, error) {

	var msg *types.Message
	var e error

	wait := sync.WaitGroup{}
	wait.Add(1)

	go func(d *DialogRouter, m *types.Message) {
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

		msgCtx, err := d.ContextStorage.Get(m.ConversationId)
		if err != nil {
			fmt.Println(err.Error())

			d.ContextStorage.Save(m.ConversationId, &types.DialogContext{
				ConversationData:          make(map[string]string, 1),
				PerUserInConversationData: make(map[string]string, 1),
				UserData:                  make(map[string]string, 1),
			})
			// TODO - Just means the ctx doesn't exist
		}

		resp, err := d.Dialogs[highestIndex].MessageReceived(msgCtx, m)
		if err != nil {
			e = err
			wait.Done()
			return
		}

		msg = resp
		e = nil
		wait.Done()
	}(router, message)

	wait.Wait()
	return msg, e
}

func getHighestScoringDialog(dialogs []types.Dialoger, msg *types.Message) (int, error) {

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