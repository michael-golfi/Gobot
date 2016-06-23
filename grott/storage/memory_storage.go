package storage

import "fmt"

type InMemoryStorage struct {
	data map[string]interface{}
}

func (s *InMemoryStorage) Get(id string, data chan interface{}) error {
	if s.data[id] != nil {
		data <- s.data[id]
		return nil
	}
	return error.Error(fmt.Sprintf("No data with id: %s found", id))
}

func (s *InMemoryStorage) Save(id string, data interface{}, err chan error) {
	s.data[id] = data
}

func (s *InMemoryStorage) Delete(id string) {
	delete(s.data, id)
}

