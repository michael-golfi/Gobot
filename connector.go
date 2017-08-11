package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/michael-golfi/Grott/request"
)

var (
	client = request.BotHTTPClient{}
)

// CreateConversation - Needs proper testing
func CreateConversation(activity *Activity, params *ConversationParameters) error {
	url := fmt.Sprintf(request.CreateConversationPath, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(params); err != nil {
		return err
	}

	if _, err := client.Post(url, b); err != nil {
		return err
	}

	return nil
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
	url := fmt.Sprintf(request.DeleteActivityPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("DeleteActivity: %s", url)

	if _, err := client.Delete(url); err != nil {
		return err
	}

	return nil
}

// GetConversationMembers - Successfully returns members
func GetConversationMembers(activity *Activity) ([]ChannelAccount, error) {
	url := fmt.Sprintf(request.GetConversationMembersPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("GetConversationMembers: %s", url)

	res, err := client.Get(url)
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
	url := fmt.Sprintf(request.GetActivityMembersPath, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("GetActivityMembers: %s", url)

	res, err := client.Get(url)
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
	url := fmt.Sprintf(request.ReplyToActivityPath, activity.ServiceURL, activity.Conversation.ID, activity.ReplyToID)
	log.Printf("ReplyToActivity: %s", url)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return err
	}

	if _, err := client.Post(url, b); err != nil {
		return err
	}

	return nil
}

// SendToConversation - Successfully sends messages
func SendToConversation(activity *Activity) error {
	url := fmt.Sprintf(request.SendToConversationPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("SendToConversation: %s", url)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return err
	}

	if _, err := client.Post(url, b); err != nil {
		return err
	}

	return nil
}

// UploadAttachment - Needs proper testing
func UploadAttachment(activity *Activity, attachment *AttachmentUpload) error {
	url := fmt.Sprintf(request.UploadAttachmentPath, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("UploadAttachment: %s", url)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return err
	}

	if _, err := client.Post(url, b); err != nil {
		return err
	}

	return nil
}
