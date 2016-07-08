package storage

import "github.com/michael-golfi/Grott/grott/types"

type ContextStorage interface {
	Get(id string) (*types.DialogContext, error)
	Save(id string, data *types.DialogContext) error
	Delete(id string)
}