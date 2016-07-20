package inmemory

import (
	"errors"
	"fmt"
	"github.com/michael-golfi/Grott/grott/storage"
)

type InMemoryStorage struct {
	cache map[string]*storage.DialogContext
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cache: make(map[string]*storage.DialogContext),
	}
}

func (s InMemoryStorage) Get(id string) (*storage.DialogContext, error) {
	fmt.Printf("InMemoryStorage: Get: %s\n", id)
	if s.cache[id] != nil {
		return s.cache[id], nil
	}

	return nil, errors.New("No data with id: " + id + " found")
}

func (s InMemoryStorage) Save(id string, data *storage.DialogContext) error {
	fmt.Printf("InMemoryStorage: Save: %s\n", id)
	s.cache[id] = data

	return nil
}

func (s InMemoryStorage) Delete(id string) {
	fmt.Printf("InMemoryStorage: Delete: %s\n", id)
	delete(s.cache, id)
}
