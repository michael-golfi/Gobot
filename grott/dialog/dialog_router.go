package dialog

import (
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/storage"
)

type DialogRouter struct {
	Dialogs        []types.Dialoger
	ContextStorage storage.Storage
}

func Start(dialogs []types.Dialoger) *DialogRouter {
	inMemoryStorage := storage.InMemoryStorage{}

	return &DialogRouter{
		Dialogs: dialogs,
		ContextStorage: inMemoryStorage,
	}
}

func StartWithStorage(dialogs []types.Dialoger, storage storage.Storage) *DialogRouter {
	return &DialogRouter{
		Dialogs: dialogs,
		ContextStorage: storage,
	}
}

func (router *DialogRouter) HandleMessage(message types.Message) (types.Message, error) {

	respChan := make(chan types.Message, 1)
	errChan := make(chan error, 1)

	go func(r chan types.Message, e chan error, d DialogRouter, m types.Message) {
		highestIndex := 0
		if len(d.Dialogs) > 1 {
			highestIndex = getHighestScoringDialog(d.Dialogs, m)
		}

		msgCtx, err := d.ContextStorage.Get(m.ConversationId)
		if err != nil {
			e <- err
			r <- nil
			return
		}

		resp, err := d.Dialogs[highestIndex].MessageReceived(msgCtx, m)
		if err != nil {
			e <- err
			r <- nil
			return
		}

		r <- resp
		e <- nil
	}(respChan, errChan, router, message)

	return <-respChan, <-errChan
}

func getHighestScoringDialog(dialogs []types.Dialoger, msg types.Message) *types.Dialoger {

	highestIndex := 0
	highestScore := 0

	for i, dialog := range dialogs {

		score, err := dialog.CalculateScore(msg)

		if err != nil {
			return err
		}

		if score > highestScore {
			highestScore = score
			highestIndex = i
		}

	}

	return highestIndex
}