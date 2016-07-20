package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
	"net/http"
	"net/url"
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
	uri, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + activity.Conversation.Id + "/activities"

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return post(uri.String(), b, headers)
}

func (c *ClientConnector) Respond(activity types.Activity, conversationId, activityId string, headers map[string]string) (*http.Response, error) {
	uri, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + activity.Conversation.Id + "/activities" + "/" + activity.Id

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(activity); err != nil {
		return nil, err
	}

	return post(uri.String(), b, headers)
}

func (c *ClientConnector) GetMembers(conversationId, activityId string, headers map[string]string) (*http.Response, error) {
	uri, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + conversationId + "/activities" + activityId + "/members"

	return get(uri.String(), headers)
}

func (c *ClientConnector) UploadAttachment(attachment types.AttachmentUpload, conversationId string, headers map[string]string) (*http.Response, error) {
	uri, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + conversationId + "/attachments"

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(attachment); err != nil {
		return nil, err
	}

	return post(uri.String(), b, headers)
}
