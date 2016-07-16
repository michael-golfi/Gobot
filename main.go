package main

import (
	"github.com/michael-golfi/Grott/grott/connector"
	"github.com/michael-golfi/Grott/grott/types"
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"sync"
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
		Id: "j23oi4j23oi4j32o4ij324",
		Timestamp: time.Now(),
		ServiceUrl: "http://localhost:9000",
		ChannelId: "emulator",
		From: types.ChannelAccount{
			Id: "2343jd",
			Name: "User1",
		},
		Conversation: types.ConversationAccount{
			Id: "3i4j23oi4",
			Name: "Conv1",
			IsGroup: false,
		},
		Recipient: types.ChannelAccount{
			Id: "34jj23j2j",
			Name: "Bot1",
		},
		Text: "sent [Ping] to bot",
		Attachments: []types.Attachment{},
		// Entities: []types.Entity{},
	}

	wait := sync.WaitGroup{}
	wait.Add(1)

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handled")
		wait.Done()
	})

	go http.ListenAndServe(":3798", nil)
	
	headers := map[string]string{}

	conn.Send(activity, "3i4j23oi4", headers)

	wait.Wait()
}
