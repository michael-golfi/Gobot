package connector_test

import (
	"encoding/json"
	"github.com/michael-golfi/Grott/grott/connector"
	"github.com/michael-golfi/Grott/grott/activity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

/**
This test requires running the Bot Emulator
*/
func TestClientConnector_CreateConversation_Send(t *testing.T) {
	url := "http://localhost:9000"
	conn := connector.NewClientConnector(url)

	user := activity.ChannelAccount{
		Id:   "2c1c7fa3",
		Name: "User1",
	}

	conversation := activity.ConversationAccount{
		Id:      "8a684db8",
		Name:    "Conv1",
		IsGroup: false,
	}

	recipient := activity.ChannelAccount{
		Id:   "Bot1",
		Name: "Bot1",
	}

	activity := activity.Activity{
		Type:         "message",
		Id:           "a8e5599a0e434e92994ebec5506d2be5",
		Timestamp:    time.Now(),
		ServiceUrl:   "http://localhost:9000/",
		ChannelId:    "emulator",
		From:         user,
		Conversation: conversation,
		Recipient:    recipient,
		Text:         "Hello User!",
		Attachments:  []activity.Attachment{},
		Entities:     []activity.Entity{},
	}

	headers := map[string]string{}

	conversationParameters := activity.ConversationParameters{
		Bot:       recipient,
		IsGroup:   false,
		Members:   []activity.ChannelAccount{user},
		TopicName: "default",
	}

	resp, err := conn.CreateConversation(conversationParameters, headers)
	assert.NoError(t, err)

	b, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	m := map[string]string{}
	err = json.Unmarshal([]byte(string(b)), &m)
	assert.NoError(t, err)

	conversation.Id = m["id"]

	resp, err = conn.Send(activity, conversation.Id, headers)
	assert.NoError(t, err)

	b, err = ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
}
