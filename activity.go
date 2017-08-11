package bot

import "time"

// Activity defines a message that is exchanged between bot and user.
type Activity struct {
	Action           string                 `json:"action,omitempty"`
	Attachments      []Attachment           `json:"attachments,omitempty"`
	AttachmentLayout string                 `json:"attachmentLayout,omitempty"`
	ChannelData      interface{}            `json:"channelData,omitempty"`
	ChannelID        string                 `json:"channelId,omitempty"`
	Conversation     ConversationAccount    `json:"conversation"`
	Entities         []interface{}          `json:"entities,omitempty"`
	From             ChannelAccount         `json:"from"`
	HistoryDisclosed bool                   `json:"historyDisclosed,omitempty"`
	ID               string                 `json:"id,omitempty"`
	InputHint        string                 `json:"inputHint,omitempty"`
	Locale           string                 `json:"locale,omitempty"`
	LocalTimestamp   time.Time              `json:"localTimestamp,omitempty"`
	MembersAdded     []ChannelAccount       `json:"membersAdded,omitempty"`
	MembersRemoved   []ChannelAccount       `json:"membersRemoved,omitempty"`
	Recipient        ChannelAccount         `json:"recipient"`
	RelatesTo        *ConversationReference `json:"relatesTo,omitempty"`
	ReplyToID        string                 `json:"replyToId,omitempty"`
	ServiceURL       string                 `json:"serviceUrl,omitempty"`
	Speak            string                 `json:"speak,omitempty"`
	SuggestedActions *SuggestedActions      `json:"suggestedActions,omitempty"`
	Summary          string                 `json:"summary,omitempty"`
	Text             string                 `json:"text,omitempty"`
	TextFormat       string                 `json:"textFormat,omitempty"`
	Timestamp        time.Time              `json:"timestamp,omitempty"`
	TopicName        string                 `json:"topicName,omitempty"`
	Type             string                 `json:"type,omitempty"`
}

// AnimationCard defines a card that can play animated GIFs or short videos.
type AnimationCard struct {
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailURL `json:"image"`
	Media     []MediaURL   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}

// Attachment defines additional information to include in the message. An attachment may be a media file (e.g., audio, video, image, file) or a rich card.
type Attachment struct {
	ContentType  string      `json:"contentType"`
	ContentURL   string      `json:"contentUrl"`
	Content      interface{} `json:"content"`
	Name         string      `json:"name"`
	ThumbnailURL string      `json:"thumbnailUrl"`
}

// AttachmentInfo describes an attachment.
type AttachmentInfo struct {
	Name  string           `json:"name"`
	Type  string           `json:"type"`
	Views []AttachmentView `json:"views"`
}

// AttachmentView defines a attachment view.
type AttachmentView struct {
	ViewID string `json:"viewId"`
	Size   int    `json:"size"`
}

// AttachmentUpload defines an attachment to be uploaded.
type AttachmentUpload struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	OriginalBase64  string `json:"originalBase64"`
	ThumbnailBase64 string `json:"thumbnailBase64"`
}

// AudioCard defines a card that can play an audio file.
type AudioCard struct {
	Aspect    string       `json:"aspect"`
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailURL `json:"image"`
	Media     []MediaURL   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}

// BotData defines state data for a user, a conversation, or a user in the context of a specific conversation that is stored using the Bot State service.
type BotData struct {
	Data interface{} `json:"data"`
	ETag string      `json:"eTag"`
}

// CardAction defines an action to perform.
type CardAction struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
	Value string `json:"value"`
}

// CardImage defines an image to display on a card.
type CardImage struct {
	Alt string     `json:"alt"`
	Tap CardAction `json:"tap"`
	URL string     `json:"url"`
}

// ChannelAccount defines a bot or user account on the channel.
type ChannelAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Conversation defines a conversation, including the bot and users that are included within the conversation.
type Conversation struct {
	Bot       ChannelAccount   `json:"bot"`
	IsGroup   bool             `json:"isGroup"`
	Members   []ChannelAccount `json:"members"`
	TopicName string           `json:"topicName"`
	Activity  Activity         `json:"activity"`
}

// ConversationAccount defines a conversation in a channel.
type ConversationAccount struct {
	ID      string `json:"id"`
	IsGroup bool   `json:"isGroup"`
	Name    string `json:"name"`
}

// ConversationReference defines a particular point in a conversation.
type ConversationReference struct {
	ActivityID   string              `json:"activityId"`
	Bot          ChannelAccount      `json:"bot"`
	ChannelID    string              `json:"channelId"`
	Conversation ConversationAccount `json:"conversation"`
	ServiceURL   string              `json:"serviceUrl"`
	User         ChannelAccount      `json:"user"`
}

// Error defines an error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse defines an HTTP API response.
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Fact defines a key-value pair that contains a fact.
type Fact struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GeoCoordinates defines a geographical location using World Geodetic System (WSG84) coordinates.
type GeoCoordinates struct {
	Elevation int     `json:"elevation"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
}

// HeroCard defines a card with a large image, title, text, and action buttons.
type HeroCard struct {
	Buttons  []CardAction `json:"buttons"`
	Images   []CardImage  `json:"images"`
	Subtitle string       `json:"subtitle"`
	Tap      CardAction   `json:"tap"`
	Text     string       `json:"text"`
	Title    string       `json:"title"`
}

// ID identifies a resource.
type ID string

// MediaURL defines the URL to a media file's source.
type MediaURL struct {
	Profile string `json:"profile"`
	URL     string `json:"url"`
}

// Mention defines a user or bot that was mentioned in the conversation.
type Mention struct {
	Mentioned ChannelAccount `json:"mentioned"`
	Text      string         `json:"text"`
	Type      string         `json:"type"`
}

// Place defines a place that was mentioned in the conversation.
type Place struct {
	Geo  GeoCoordinates `json:"geo"`
	Name string         `json:"name"`
	Type string         `json:"type"`
}

// ReceiptCard defines a card that contains a receipt for a purchase.
type ReceiptCard struct {
	Buttons []CardAction  `json:"buttons"`
	Facts   []Fact        `json:"facts"`
	Items   []ReceiptItem `json:"items"`
	Tap     CardAction    `json:"tap"`
	Tax     string        `json:"tax"`
	Title   string        `json:"title"`
	Total   string        `json:"total"`
	Vat     string        `json:"vat"`
}

// ReceiptItem defines a line item within a receipt.
type ReceiptItem struct {
	Image    CardImage  `json:"image"`
	Price    string     `json:"price"`
	Quantity string     `json:"quantity"`
	Subtitle string     `json:"subtitle"`
	Tap      CardAction `json:"tap"`
	Text     string     `json:"text"`
	Title    string     `json:"title"`
}

// ResourceResponse defines a resource.
type ResourceResponse struct {
	ActivityID string `json:"activityId"`
	ServiceURL string `json:"serviceUrl"`
	ID         string `json:"id"`
}

// SignInCard defines a card that lets a user sign in to a service.
type SignInCard struct {
	Buttons []CardAction `json:"buttons"`
	Text    string       `json:"text"`
}

// SuggestedActions defines the options from which a user can choose.
type SuggestedActions struct {
	To      []string     `json:"to"`
	Actions []CardAction `json:"actions"`
}

// ThumbnailCard defines a card with a thumbnail image, title, text, and action buttons.
type ThumbnailCard struct {
	Buttons  []CardAction `json:"buttons"`
	Images   []CardImage  `json:"images"`
	Subtitle string       `json:"subtitle"`
	Tap      CardAction   `json:"tap"`
	Text     string       `json:"text"`
	Title    string       `json:"title"`
}

// ThumbnailURL defines the URL to an image's source.
type ThumbnailURL struct {
	Alt string `json:"alt"`
	URL string `json:"url"`
}

// VideoCard defines a card that can play videos.
type VideoCard struct {
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailURL `json:"image"`
	Media     []MediaURL   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}

/* END Main Activity graph types */

// Activity Types
var (
	Message = "message"
)

func (a *Activity) CreateReply(text, locale string) Activity {
	return Activity{
		//Action:           "",
		Attachments: nil,
		//AttachmentLayout: "",
		ChannelData:  nil,
		ChannelID:    a.ChannelID,
		Conversation: a.Conversation,
		Entities:     []interface{}{},
		From:         a.Recipient,
		//HistoryDisclosed: false,
		//ID: nil,
		LocalTimestamp: time.Now(),
		Type:           Message,
		Timestamp:      time.Now().UTC(),

		Recipient:  a.From,
		ReplyToID:  a.ID,
		ServiceURL: a.ServiceURL,

		Text:       text,
		TextFormat: "plain",
		Locale:     locale,

		//MembersAdded:   []ChannelAccount{},
		//MembersRemoved: []ChannelAccount{},
		RelatesTo: ConversationReference{
			ActivityID:   a.ID,
			Bot:          a.Recipient,
			ChannelID:    a.ChannelID,
			Conversation: a.Conversation,
			ServiceURL:   a.ServiceURL,
			User:         a.From,
		},
	}
}
