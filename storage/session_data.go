package storage

type SessionData struct {
	syncMap Map
}

func (m *SessionData) Load(key string) (value string, ok bool) {
	val, ok := m.syncMap.Load(key)

	value, okConv := val.(string)
	if !okConv {
		return "", false
	}

	return value, ok
}

func (m *SessionData) Store(key, value string) {
	m.syncMap.Store(key, value)
}

func (m *SessionData) Delete(key string) {
	m.syncMap.Delete(key)
}
