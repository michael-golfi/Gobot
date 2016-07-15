package storage

import (
	"errors"
	"github.com/michael-golfi/Grott/grott/types"
	"fmt"
)

type InMemoryStorage struct {
	cache map[string]*types.DialogContext
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cache: make(map[string]*types.DialogContext),
	}
}

func (s InMemoryStorage) Get(id string) (*types.DialogContext, error) {
	fmt.Printf("InMemoryStorage: Get: %s\n", id)
	if s.cache[id] != nil {
		return s.cache[id], nil
	}

	return nil, errors.New("No data with id: " + id + " found")
}

func (s InMemoryStorage) Save(id string, data *types.DialogContext) error {
	fmt.Printf("InMemoryStorage: Save: %s\n", id)
	s.cache[id] = data

	return nil
}

func (s InMemoryStorage) Delete(id string) {
	fmt.Printf("InMemoryStorage: Delete: %s\n", id)
	delete(s.cache, id)
}