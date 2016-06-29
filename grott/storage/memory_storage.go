package storage

import (
	"errors"
)

type InMemoryStorage struct {
	store map[string]string
	//cache map[types.BotDataKey]types.BotDataEntry
	cache map[string]interface{}
}

func (s *InMemoryStorage) Get(id string) (interface{}, error) {
	if s.cache[id] != nil {
		return &s.cache[id], nil
	}

	return nil, errors.New("No data with id: " + id + " found")
}

func (s *InMemoryStorage) Save(id string, data interface{}) error {
	s.cache[id] = data
	return nil
}

func (s *InMemoryStorage) Delete(id string) {
	delete(s.cache, id)
}