package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	Create    = "%s/v3/conversations"
	Send      = "%s/v3/conversations/%s/activities"
	Reply     = "%s/v3/conversations/%s/activities/%s"
	Upload    = "%s/v3/conversations/%s/attachments"
	GetMember = "%s/v3/conversations/%s/activities/%s/members"
)

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

func CreateConversation(activity *Activity, params *ConversationParameters) (string, error) {
	url := fmt.Sprintf(Create, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)
	_, err := postParams(url, params)
	return "", err
}

// TODO
func CreateDirectConversation(activity *Activity) error {
	url := fmt.Sprintf(Create, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)
	_, err := post(url, activity)
	return err
}

// TODO
func DeleteActivity(activity *Activity) error {
	url := fmt.Sprintf(Create, activity.ServiceURL)
	log.Printf("Createconversation: %s", url)
	_, err := post(url, activity)
	return err
}

// TODO
func GetActivityMembers(activity *Activity) (*http.Response, error) {
	url := fmt.Sprintf(GetMember, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("GetMembers: %s", url)
	return http.Get(url)
}

// TODO
func GetConversationMembers(activity *Activity) (*http.Response, error) {
	url := fmt.Sprintf(GetMember, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("GetMembers: %s", url)
	return http.Get(url)
}

func ReplyToActivity(activity *Activity) error {
	url := fmt.Sprintf(Reply, activity.ServiceURL, activity.Conversation.ID, activity.ID)
	log.Printf("Respond: %s", url)
	_, err := post(url, activity)
	return err
}

func SendToConversation(activity *Activity) error {
	url := fmt.Sprintf(Send, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("SendMessage: %s", url)
	_, err := post(url, activity)
	return err
}

func UploadAttachment(activity *Activity, attachment *AttachmentUpload) error {
	url := fmt.Sprintf(Upload, activity.ServiceURL, activity.Conversation.ID)
	log.Printf("UploadAttachment: %s", url)
	_, err := post(url, activity)
	return err
}

func post(url string, activity *Activity) (*http.Response, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return http.Post(url, "application/json", b)
}

func postParams(url string, params *ConversationParameters) (*http.Response, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return http.Post(url, "application/json", b)
}
