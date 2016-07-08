package types

type Message struct {
	MessageType           string `json:"type"`
	Id                    string `json:"id"`
	ConversationId        string `json:"conversationId"`
	Created               string `json:"created"`
	SourceText            string `json:"sourceText"`
	SourceLanguage        string `json:"sourceLanguage"`
	Language              string `json:"language"`
	Text                  string `json:"text"`
	Attachments           []Attachment `json:"attachments"`
	From                  Account `json:"from"`
	ReplyToMessageId      string `json:"replyToMessageId"`
	Participants          []Account `json:"participants"`
	Mentions              []Mention `json:"mentions"`
	Place                 string `json:"place"`
	ChannelMessageId      string `json:"channelMessageId"`
	ChannelConversationId string `json:"channelConversationId"`
	ChannelData           interface{} `json:"channelData"`
	Location              Location `json:"location"`
	Hashtags              []string `json:"hashtags"`
	ETag                  string `json:"ETag"`
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
	Actions      []MessageAction `json:"actions"`
	ContentType  string `json:"contentType"`
	ContentUrl   string `json:"contentUrl"`
	FallbackText string `json:"fallbackText"`
	Title        string `json:"title"`
	TitleLink    string `json:"titleLink"`
	Text         string `json:"text"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type MessageAction struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Image   string `json:"image"`
}

type Account struct {
	Name      string `json:"name"`
	ChannelId string `json:"channelId"`
	Address   string `json:"address"`
	Id        string `json:"id"`
	IsBot     bool `json:"isBot"`
}

type Mention struct {
	Mentioned Account `json:"mentioned"`
	Text      string `json:"text"`
}

type Location struct {
	Altitude  float32 `json:"altitude"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Localizer interface {
	GetText(language, messageId string) string
	NGetText(language, messageId, messageIdPlural string, count int) string
}

type SessionOptions struct {
	Dialogs      []Dialoger `json:"dialogs"`
	DialogId     string `json:"dialogId"`
	DialogArgs   interface{} `json:"dialogArgs"`
	Localizer    Localizer `json:"localizer"`
	MinSendDelay int32 `json:"minSendDelay"`
}

type BotConnectorMessage struct {
	Message `json:"message"`
	BotUserData                  map[string]interface{} `json:"botUserData"`
	BotConversationData          map[string]interface{} `json:"botConversationData"`
	BotPerUserInConversationData map[string]interface{} `json:"botPerUserInConversationData"`
}