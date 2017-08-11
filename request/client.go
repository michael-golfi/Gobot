package request

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
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

type BotHTTPClient struct {
}

type BotHTTPConfig struct {
	CreateConversation     string
	SendToConversation     string
	UpdateActivity         string
	DeleteActivity         string
	ReplyToActivity        string
	GetConversationMembers string
	GetActivityMembers     string
	UploadAttachment       string
}

func (b *BotHTTPClient) Post(url string, body io.Reader) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	spew.Dump(req)
	return client.Do(req)
}

func (b *BotHTTPClient) Get(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	return client.Do(req)
}

func (b *BotHTTPClient) Put(url string, body io.Reader) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	return client.Do(req)
}

func (b *BotHTTPClient) Delete(url string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	return client.Do(req)
}
