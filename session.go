package bot

import "github.com/michael-golfi/Grott/storage"

// All globally accessible data needs to be thread-safe
// Dialog data is private to each interaction so it is thread-private

type Session interface {
	GetUserData() (*storage.SessionData, error)
	GetConversationData() (*storage.SessionData, error)
	GetPrivateConversationData() (*storage.SessionData, error)

	SetUserData() error
	SetConversationData() error
	SetPrivateConversationData() error

	ClearUserState()
}
