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

func (r *DialogRouter) HandleMessage(msg types.Message) error {

	highestIndex := getHighestScoringDialog(r.Dialogs, msg)

	msgCtx,err := r.ContextStorage.Get(msg.ConversationId)
	if err != nil {
		return err
	}

	r.Dialogs[highestIndex].MessageReceived(msgCtx, msg)

	return nil
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