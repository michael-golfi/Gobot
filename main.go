package main

import (
	"github.com/michael-golfi/Grott/grott/connector"
	"github.com/michael-golfi/Grott/grott/types"
	"time"
)

func main() {
	url := "http://localhost:9000"
	conn := connector.NewClientConnector(url)
	/*
		bot := types.ChannelAccount{
			Id: "Bot1",
			Name: "Test",
		}

		me := types.ChannelAccount{
			Id: "2",
			Name: "User1",
		}*/

	activity := types.Activity{
		Type: "ping",
		Id: "a8e5599a0e434e92994ebec5506d2be5",
		Timestamp: time.Now(),
		ServiceUrl: "http://localhost:9000",
		ChannelId: "emulator",
		From: types.ChannelAccount{
			Id: "2c1c7fa3",
			Name: "User1",
		},
		Conversation: types.ConversationAccount{
			Id: "8a684db8",
			Name: "Conv1",
			IsGroup: false,
		},
		Recipient: types.ChannelAccount{
			Id: "56800324",
			Name: "Bot1",
		},
		Text: "sent [Ping] to user",
		Attachments: []types.Attachment{},
		// Entities: []types.Entity{},
	}

	headers := map[string]string{}

	conn.Send(activity, "8a684db8", headers)
}
