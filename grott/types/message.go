package types

type Message struct {
	MessageType           string
	Id                    string
	ConversationId        string
	Created               string
	SourceText            string
	SourceLanguage        string
	Language              string
	Text                  string
	Attachments           []Attachment
	From                  Account
	ReplyToMessageId      string
	Participants          []Account
	Mentions              []Mention
	Place                 string
	ChannelMessageId      string
	ChannelConversationId string
	ChannelData           interface{}
	Location              Location
	Hashtags              []string
	ETag                  string
}

type IMessage interface {
	SetLanguage(language string)
	//SetText(session Session, prompt []string, args interface{})
	//SetNText(Session Session, msg string, msgPlural string, count int)
	//ComposePrompt(Session Session, prompts [][]string, args interface{})
	AddAttachment(attachment Attachment)
	SetChannelData(data interface{})
}

type Attachment struct {
	Actions      []Action
	ContentType  string
	ContentUrl   string
	FallbackText string
	Title        string
	TitleLink    string
	Text         string
	ThumbnailUrl string
}

type Action struct {
	Title   string
	Message string
	Url     string
	Image   string
}

type Account struct {
	Name      string
	ChannelId string
	Address   string
	Id        string
	IsBot     bool
}

type Mention struct {
	Mentioned Account
	Text      string
}

type Location struct {
	Altitude  float32
	Latitude  float32
	Longitude float32
}

type Localizer interface {
	GetText(language, messageId string) string
	NGetText(language, messageId, messageIdPlural string, count int) string
}

type SessionOptions struct {
	Dialogs      []Dialoger
	DialogId     string
	DialogArgs   interface{}
	Localizer    Localizer
	MinSendDelay int32
}

type BotConnectorMessage struct {
	Message
	BotUserData                  map[string]interface{}
	BotConversationData          map[string]interface{}
	BotPerUserInConversationData map[string]interface{}
}
