package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ref: https://github.com/Microsoft/BotBuilder/blob/master/CSharp/Library/Microsoft.Bot.Connector.Shared/ConnectorAPI/Conversations.cs
var (
	CreateConversationPath     = "%s/v3/conversations"
	SendToConversationPath     = "%s/v3/conversations/%s/activities"
	UpdateActivityPath         = "%s/v3/conversations/%s/activities/%s"
	DeleteActivityPath         = "%s/v3/conversations/%s/activities"
	ReplyToActivityPath        = "%s/v3/conversations/%s/activities/%s"
	GetConversationMembersPath = "%s/v3/conversations/%s/members"
	GetActivityMembersPath     = "%s/v3/conversations/%s/activities/%s/members"
	UploadAttachmentPath       = "%s/v3/conversations/%s/attachments"
)

// CreateConversation - Needs proper testing
func CreateConversation(activity *Activity, params *ConversationParameters) error {
	url := fmt.Sprintf(CreateConversationPath, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)
	_, err := postParams(url, params)
	return err
}

/*
// TODO
func CreateDirectConversation(activity *Activity) error {
	url := fmt.Sprintf(CreateDire, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)
	_, err := post(url, activity)
	return err
}*/

// DeleteActivity - Needs proper testing
func DeleteActivity(activity *Activity) error {
	url := fmt.Sprintf(DeleteActivityPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("DeleteActivity: %s", url)
	_, err := post(url, activity)
	return err
}

// GetConversationMembers - Successfully returns members
func GetConversationMembers(activity *Activity) ([]ChannelAccount, error) {
	url := fmt.Sprintf(GetConversationMembersPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("GetConversationMembers: %s", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var members []ChannelAccount
	if err := json.NewDecoder(res.Body).Decode(&members); err != nil {
		return nil, err
	}
	return members, nil
}

// GetActivityMembers - Successfully returns members
func GetActivityMembers(activity *Activity) ([]ChannelAccount, error) {
	url := fmt.Sprintf(GetActivityMembersPath, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("GetActivityMembers: %s", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var members []ChannelAccount
	if err := json.NewDecoder(res.Body).Decode(&members); err != nil {
		return nil, err
	}
	return members, nil
}

// ReplyToActivity - Successfully sends messages
func ReplyToActivity(activity *Activity) error {
	url := fmt.Sprintf(ReplyToActivityPath, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("ReplyToActivity: %s", url)
	_, err := post(url, activity)
	return err
}

// SendToConversation - Successfully sends messages
func SendToConversation(activity *Activity) error {
	url := fmt.Sprintf(SendToConversationPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("SendToConversation: %s", url)
	_, err := post(url, activity)
	return err
}

// UploadAttachment - Needs proper testing
func UploadAttachment(activity *Activity, attachment *AttachmentUpload) error {
	url := fmt.Sprintf(UploadAttachmentPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("UploadAttachment: %s", url)
	_, err := post(url, activity)
	return err
}
