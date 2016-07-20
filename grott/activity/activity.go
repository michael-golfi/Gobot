package activity

import "time"

type Activity struct {
	Type             string              `json:"type"`
	Id               string              `json:"id"`
	ServiceUrl       string              `json:"serviceUrl"`
	Timestamp        time.Time           `json:"timestamp"`
	ChannelId        string              `json:"channelId"`
	From             ChannelAccount      `json:"from"`
	Conversation     ConversationAccount `json:"conversation"`
	Recipient        ChannelAccount      `json:"recipient"`
	Locale           string       `json:"locale"`
	Text             string       `json:"text"`
	Summary          string       `json:"summary"`
	TextFormat       string       `json:"textFormat"`
	AttachmentLayout string       `json:"attachmentLayout"`
	Attachments      []Attachment `json:"attachments"`
	Entities         []Entity     `json:"entities"`
	ChannelData      interface{}  `json:"channelData"`
	ReplyToId        string       `json:"replyToId"`
	Mentions         []Mention    `json:"mentions"`
}

type Entity struct {
	Type string `json:"type"`
}

type ApiResponse struct {
	Message string `json:"message"`
}

type ChannelAccount struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ConversationAccount struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IsGroup bool   `json:"isGroup"`
}

type ConversationParameters struct {
	Bot       ChannelAccount   `json:"bot"`
	IsGroup   bool             `json:"isGroup"`
	Members   []ChannelAccount `json:"members"`
	TopicName string           `json:"topicName"`
}

type AttachmentUpload struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	OriginalBase64  []byte `json:"originalBase64"`
	ThumbnailBase64 []byte `json:"thumbnailBase64"`
}

func (a Activity) CreateReply(msg string) Activity {
	activity := a
	activity.From = a.Recipient
	activity.Recipient = a.From
	return activity
}

type Attachment struct {
	Actions      []MessageAction `json:"actions"`
	ContentType  string          `json:"contentType"`
	ContentUrl   string          `json:"contentUrl"`
	FallbackText string          `json:"fallbackText"`
	Title        string          `json:"title"`
	TitleLink    string          `json:"titleLink"`
	Text         string          `json:"text"`
	ThumbnailUrl string          `json:"thumbnailUrl"`
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
	IsBot     bool   `json:"isBot"`
}

type Mention struct {
	Mentioned Account `json:"mentioned"`
	Text      string  `json:"text"`
}