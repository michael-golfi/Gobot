package types

import "time"

type Activity struct {
	Type             string              `json:"type"`
	Id               string              `json:"id"`
	ServiceUrl       string              `json:"serviceUrl"`
	Timestamp        time.Time           `json:"timestamp"`
	ChannelId        string              `json:"channelId"`
	From             ChannelAccount      `json:"from"`
	Conversation     ConversationAccount `json:"conversationAccount"`
	Recipient        ChannelAccount      `json:"recipient"`

	Locale           string `json:"locale"`
	Text             string `json:"text"`
	Summary          string `json:"summary"`
	TextFormat       string `json:"textFormat"`
	AttachmentLayout string `json:"attachmentLayout"`
	Attachments      []Attachment `json:"attachments"`
	Entities         []Entity `json:"entities"`
	ChannelData      interface{} `json:"channelData"`
	ReplyToId        string `json:"replyToId"`
	Mentions         []Mention `json:"mentions"`
}

type Entity struct {

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

func NewConversationParameters(bot ChannelAccount, isGroup bool, members []ChannelAccount, topic string) *ConversationParameters {
	return &ConversationParameters{
		Bot: bot,
		IsGroup: isGroup,
		Members: members,
		TopicName: topic,
	}
}

func NewAttachmentUpload(typeName, name string, originalBase64, thumbnailBase64 []byte) *AttachmentUpload {
	return &AttachmentUpload{
		Type: typeName,
		Name: name,
		OriginalBase64: originalBase64,
		ThumbnailBase64: thumbnailBase64,
	}
}

func (a Activity) CreateReply(msg string) Activity {
	activity := a
	activity.From = a.Recipient
	activity.Recipient = a.From
	return activity
}