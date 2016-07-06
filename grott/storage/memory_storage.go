package storage

import (
	"errors"
	"github.com/michael-golfi/Grott/grott/types"
)

type InMemoryStorage struct {
	cache map[string]*types.DialogContext
}

func (s InMemoryStorage) Get(id string) (*types.DialogContext, error) {
	if s.cache[id] != nil {
		return s.cache[id], nil
	}

	return nil, errors.New("No data with id: " + id + " found")
}

func (s InMemoryStorage) Save(id string, data types.DialogContext) error {
	s.cache[id] = &data
	return nil
}

func (s InMemoryStorage) Delete(id string) {
	delete(s.cache, id)
}