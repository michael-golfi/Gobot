package storage

import "errors"

type InMemoryStorage struct {
	data map[string]interface{}
}

func (s *InMemoryStorage) Get(id string, data chan interface{}) error {
	if s.data[id] != nil {
		data <- s.data[id]
		return nil
	}
	return errors.New("No data with id: " + id + " found")
}

func (s *InMemoryStorage) Save(id string, data interface{}, err chan error) {
	s.data[id] = data
}

func (s *InMemoryStorage) Delete(id string) {
	delete(s.data, id)
}