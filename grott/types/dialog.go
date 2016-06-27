package types

type Dialog interface {
	SetInput(in chan Message)
	SetOutput(out chan Message)
	Begin(session Session, args interface{})
	ReplyReceived(session Session)
	DialogResumed(session Session, result chan interface{})
	CompareConfidence(action SessionAction, language, utterance string, score int)
}

type DialogAction interface {
}