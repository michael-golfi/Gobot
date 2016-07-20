package connector

import (
	"bytes"
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
	"net/http"
	"net/url"
)

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

	b, err := parameters.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return post(baseUrl, bytes.NewBuffer(b), headers)
}

func (c *ClientConnector) Send(activity types.Activity, conversationId string, headers map[string]string) (*http.Response, error) {
	uri, err := url.Parse(activity.ServiceUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + activity.Conversation.Id + "/activities"

	b, err := activity.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return post(uri.String(), bytes.NewBuffer(b), headers)
}

func (c *ClientConnector) Respond(activity types.Activity, conversationId, activityId string, headers map[string]string) (*http.Response, error) {
	uri, err := url.Parse(activity.ServiceUrl)
	if err != nil {
		return nil, err
	}

	uri.Path += "v3/conversations/" + activity.Conversation.Id + "/activities" + "/" + activity.Id

	b, err := activity.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return post(uri.String(), bytes.NewBuffer(b), headers)
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

	b, err := attachment.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return post(uri.String(), bytes.NewBuffer(b), headers)
}
