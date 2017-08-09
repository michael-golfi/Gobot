package bot

import "time"

type Activity struct {
	Action           string
	Attachments      []Attachment          `json:"attachments"`
	AttachmentLayout string                `json:"attachmentLayout"`
	ChannelData      interface{}           `json:"channelData"`
	ChannelID        string                `json:"channelId"`
	Conversation     ConversationAccount   `json:"conversation"`
	Entities         []interface{}         `json:"entities"`
	From             ChannelAccount        `json:"from"`
	HistoryDisclosed bool                  `json:"historyDisclosed"`
	ID               string                `json:"id"`
	InputHint        string                `json:"inputHint"`
	Locale           string                `json:"locale"`
	LocalTimestamp   time.Time             `json:"localTimestamp"`
	MembersAdded     []ChannelAccount      `json:"membersAdded"`
	MembersRemoved   []ChannelAccount      `json:"membersRemoved"`
	Recipient        ChannelAccount        `json:"recipient"`
	RelatesTo        ConversationReference `json:"relatesTo"`
	ReplyToID        string                `json:"replyToId"`
	ServiceURL       string                `json:"serviceUrl"`
	Speak            string                `json:"speak"`
	SuggestedActions SuggestedActions      `json:"suggestedActions"`
	Summary          string                `json:"summary"`
	Text             string                `json:"text"`
	TextFormat       string                `json:"textFormat"`
	Timestamp        time.Time             `json:"timestamp"`
	TopicName        string                `json:"topicName"`
	Type             string                `json:"type"`
}

type AnimationCard struct {
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailUrl `json:"image"`
	Media     []MediaUrl   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}

type Attachment struct {
	ContentType  string      `json:"contentType"`
	ContentURL   string      `json:"contentUrl"`
	Content      interface{} `json:"content"`
	Name         string      `json:"name"`
	ThumbnailURL string      `json:"thumbnailUrl"`
}

type AttachmentInfo struct {
	Name  string           `json:"name"`
	Type  string           `json:"type"`
	Views []AttachmentView `json:"views"`
}

type AttachmentView struct {
	ViewID string `json:"viewId"`
	Size   int    `json:"size"`
}

type AttachmentUpload struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	OriginalBase64  string `json:"originalBase64"`
	ThumbnailBase64 string `json:"thumbnailBase64"`
}

type AudioCard struct {
	Aspect    string       `json:"aspect"`
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailUrl `json:"image"`
	Media     []MediaUrl   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}

type BotData struct {
	Data interface{} `json:"data"`
	ETag string      `json:"eTag"`
}

type CardAction struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
	Value string `json:"value"`
}

type CardImage struct {
	Alt string     `json:"alt"`
	Tap CardAction `json:"tap"`
	URL string     `json:"url"`
}

type ChannelAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Conversation struct {
	Bot       ChannelAccount   `json:"bot"`
	IsGroup   bool             `json:"isGroup"`
	Members   []ChannelAccount `json:"members"`
	TopicName string           `json:"topicName"`
	Activity  Activity         `json:"activity"`
}

type ConversationAccount struct {
	ID      string `json:"id"`
	IsGroup bool   `json:"isGroup"`
	Name    string `json:"name"`
}

type ConversationReference struct {
	ActivityID   string              `json:"activityId"`
	Bot          ChannelAccount      `json:"bot"`
	ChannelID    string              `json:"channelId"`
	Conversation ConversationAccount `json:"conversation"`
	ServiceURL   string              `json:"serviceUrl"`
	User         ChannelAccount      `json:"user"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Fact struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GeoCoordinates struct {
	Elevation int     `json:"elevation"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
}

type HeroCard struct {
	Buttons  []CardAction `json:"buttons"`
	Images   []CardImage  `json:"images"`
	Subtitle string       `json:"subtitle"`
	Tap      CardAction   `json:"tap"`
	Text     string       `json:"text"`
	Title    string       `json:"title"`
}

type ID string

type MediaUrl struct {
	Profile string `json:"profile"`
	URL     string `json:"url"`
}

type Mention struct {
	Mentioned ChannelAccount `json:"mentioned"`
	Text      string         `json:"text"`
	Type      string         `json:"type"`
}

type Place struct {
	Geo  GeoCoordinates `json:"geo"`
	Name string         `json:"name"`
	Type string         `json:"type"`
}

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

type ReceiptItem struct {
	Image    CardImage  `json:"image"`
	Price    string     `json:"price"`
	Quantity string     `json:"quantity"`
	Subtitle string     `json:"subtitle"`
	Tap      CardAction `json:"tap"`
	Text     string     `json:"text"`
	Title    string     `json:"title"`
}

type ResourceResponse struct {
	ActivityId string `json:"activityId"`
	ServiceUrl string `json:"serviceUrl"`
	ID         string `json:"id"`
}

type SignInCard struct {
	Buttons []CardAction `json:"buttons"`
	Text    string       `json:"text"`
}

type SuggestedActions struct {
	To      []string     `json:"to"`
	Actions []CardAction `json:"actions"`
}

type ThumbnailCard struct {
	Buttons  []CardAction `json:"buttons"`
	Images   []CardImage  `json:"images"`
	Subtitle string       `json:"subtitle"`
	Tap      CardAction   `json:"tap"`
	Text     string       `json:"text"`
	Title    string       `json:"title"`
}

type ThumbnailUrl struct {
	Alt string `json:"alt"`
	URL string `json:"url"`
}

type VideoCard struct {
	Autoloop  bool         `json:"autoloop"`
	Autostart bool         `json:"autostart"`
	Buttons   []CardAction `json:"buttons"`
	Image     ThumbnailUrl `json:"image"`
	Media     []MediaUrl   `json:"media"`
	Shareable bool         `json:"shareable"`
	Subtitle  string       `json:"subtitle"`
	Text      string       `json:"text"`
	Title     string       `json:"title"`
}
