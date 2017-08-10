package bot

type ConversationParameters struct {
	Bot       ChannelAccount   `json:"bot"`
	IsGroup   bool             `json:"isGroup"`
	Members   []ChannelAccount `json:"members"`
	TopicName string           `json:"topicName"`
}

type BotConnector interface {
	CreateConversation(activity *Activity, params *ConversationParameters) error
	CreateDirectConversation(activity *Activity) error
	DeleteActivity(activity *Activity) error
	GetActivityMembers(activity *Activity) error
	GetConversationMembers(activity *Activity) error
	ReplyToActivity(activity *Activity) error
	SendToConversation(activity *Activity) error
	UpdateActivity(activity *Activity) error
	UploadAttachment(activity *Activity) error
}
