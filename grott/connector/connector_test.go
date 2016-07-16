package connector_test

import (
	"testing"
	"github.com/michael-golfi/Grott/grott/connector"
	"github.com/michael-golfi/Grott/grott/types"
)

func TestClientConnector_CreateConversation(t *testing.T) {
	url := "http://localhost:9000"
	conn := connector.NewClientConnector(url)
	//conversationParams := types.NewConversationParameters(nil, false, nil, "")
	headers := map[string]string{}

	conn.CreateConversation(types.ConversationParameters{}, headers)
}

func TestClientConnector_Send(t *testing.T) {
	url := "http://localhost:9000"
	conn := connector.NewClientConnector(url)

	activity := types.Activity{}
	conversationId := ""
	headers := map[string]string{}

	conn.Send(activity, conversationId, headers)
}