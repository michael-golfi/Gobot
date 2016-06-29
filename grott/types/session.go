package types

/*
type Session struct {
	Id           string
	Options      SessionOptions
	MessageSent  bool
	isReset      bool
	LastSentTime time.Time
	SendQueue    chan Message
	Dialogs      []*Dialoger
}

type SessionKeeper interface {
	Send(msg Message)
	Dispatch(state SessionState, msg Message)
	GetMessageReceived()
	SendMessage(msg Message)
	MessageSent()
	BeginDialog(id string, args interface{})
	ReplaceDialog(id string, args interface{})
	EndDialog(result string)
	CompareConfidence(language, utterance string, score int) // Supposed to have callback too
	Reset(dialogId string, args interface{})
	CreateMessage(text string, args interface{})
	GetSession() Session
	//ValidateCallstack()
}

type SessionAction struct {

}

type SessionState struct {

}

func NewSession() (*Session, error) {
	id, err := newUUID()
	if err != nil {
		return nil, err
	}

	return &Session{
		Id: id,
	}, nil
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8] &^ 0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6] &^ 0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
*/