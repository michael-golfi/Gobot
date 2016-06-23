package grott

type Storage interface {
	Get(id string, callback func(err error, data interface{}))
	Save(id string, interface{}, func(err error))
}
