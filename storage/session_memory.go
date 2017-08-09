package storage

type SessionMemory struct {
	UserData                *SessionData
	PrivateConversationData *SessionData
	ConversationData        *SessionData
}

func (s SessionMemory) GetUserData() (*SessionData, error) {
	return nil, nil
}

func (s SessionMemory) GetConversationData() (*SessionData, error) {
	return nil, nil
}

func (s SessionMemory) GetPrivateConversationData() (*SessionData, error) {
	return nil, nil
}

func (s SessionMemory) SetUserData() error {
	return nil
}

func (s SessionMemory) SetConversationData() error {
	return nil
}

func (s SessionMemory) SetPrivateConversationData() error {
	return nil
}

func (s SessionMemory) ClearUserState() {

}
