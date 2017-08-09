package bot

// All globally accessible data needs to be thread-safe
// Dialog data is private to each interaction so it is thread-private

type Session interface {
	GetUserData() (*SessionData, error)
	GetConversationData() (*SessionData, error)
	GetPrivateConversationData() (*SessionData, error)

	SetUserData() error
	SetConversationData() error
	SetPrivateConversationData() error

	ClearUserState()
}
