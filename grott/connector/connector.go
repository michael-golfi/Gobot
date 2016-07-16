package connector

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/michael-golfi/Grott/grott/types"
)

/**
	Implements the Bot Framework API connection to allow Bot -> User communication
 */

type ClientConnector struct {
	baseUrl string
}

func NewClientConnector(url string) *ClientConnector {
	return &ClientConnector{
		baseUrl: url,
	}
}

func (c *ClientConnector) CreateConversation(parameters types.ConversationParameters, headers map[string]string) (*http.Response, error) {
	baseUrl := fmt.Sprintf("%s/v3/conversations", c.baseUrl)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(parameters); err != nil {
		return nil, err
	}

	return post(baseUrl, b, headers)
}

func (c *ClientConnector) Send(activity types.Activity, conversationId string, headers map[string]string) (*http.Response, error) {
	baseUrl := fmt.Sprintf("%s/v3/conversations/%s/activities", c.baseUrl, conversationId)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return post(baseUrl, b, headers)
}

func (c *ClientConnector) Respond(activity types.Activity, conversationId, activityId string, headers map[string]string) (*http.Response, error) {
	baseUrl := fmt.Sprintf("%s/v3/conversations/%s/activities/%s", c.baseUrl, conversationId, activityId)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return post(baseUrl, b, headers)
}

func (c *ClientConnector) GetMembers(conversationId, activityId string, headers map[string]string) (*http.Response, error) {
	baseUrl := fmt.Sprintf("%s/v3/conversations/%s/activities/%s/members", c.baseUrl, conversationId, activityId)

	return get(baseUrl, headers)
}

func (c *ClientConnector) UploadAttachment(attachment types.AttachmentUpload, conversationId string, headers map[string]string) (*http.Response, error) {
	baseUrl := fmt.Sprintf("%s/v3/conversations/%s/attachments", c.baseUrl, conversationId)

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(attachment); err != nil {
		return nil, err
	}

	return post(baseUrl, b, headers)
}