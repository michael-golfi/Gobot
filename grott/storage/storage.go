package storage

type Storage interface {
	Get(id string, data chan interface{})
	Save(id string, data interface{}, err chan error)
	Delete(id string)
}