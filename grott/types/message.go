package types

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
