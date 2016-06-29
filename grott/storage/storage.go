package storage

type Storage interface {
	Get(id string) (interface{}, error)
	Save(id string, data interface{}) error
	Delete(id string)
}