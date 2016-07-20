package activity

import (
	"testing"
	"time"
	"encoding/json"
)

func BenchmarkActivity_MarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		message.MarshalJSON()
	}
}

func BenchmarkActivity_MarshalJSON2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(message)
	}
}

func BenchmarkActivity_UnmarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		byt := []byte(messageJson)
		message.UnmarshalJSON(byt)
	}
}

func BenchmarkActivity_UnmarshalJSON2(b *testing.B) {
	byt := []byte(messageJson)
	for n := 0; n < b.N; n++ {
		json.Unmarshal(byt, message)
	}
}

func BenchmarkConversationParameters_MarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		conversationParams.MarshalJSON()
	}
}

func BenchmarkConversationParameters_UnmarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		byt := []byte(conversationParamsJson)
		conversationParams.UnmarshalJSON(byt)
	}
}

func BenchmarkAttachmentUpload_MarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		attachmentUpload.MarshalJSON()
	}
}

func BenchmarkAttachmentUpload_UnmarshalJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		byt := []byte(attachmentUploadJson)
		attachmentUpload.UnmarshalJSON(byt)
	}
}

var (
	message = Activity{
		Type: "message",
		Id: "419f6441127744b9ab7fd7cec8c2968c",
		ServiceUrl: "http:\\\\localhost:9000",
		Timestamp: time.Now(),
		ChannelId: "emulator",
		From: ChannelAccount{
			Id: "5e4f5dfa",
			Name: "Bot1",
		},
		Conversation: ConversationAccount{
			Id: "1cf91be5",
			Name: "Conv1",
			IsGroup: false,
		},
		Recipient: ChannelAccount{
			Id: "617d3bf8",
			Name: "User1",
		},
		Locale: "",
		Text: "Hello User!",
		Summary: "",
		TextFormat: "",
		AttachmentLayout: "",
		Attachments: make([]Attachment, 1),
		Entities: make([]Entity, 1),
		ChannelData: ChannelAccount{

		},
		ReplyToId: "",
		Mentions: make([]Mention, 1),
	}

	messageJson =
	`{
	  "type": "message",
	  "id": "419f6441127744b9ab7fd7cec8c2968c",
	  "serviceUrl": "http:\/\/localhost:9000\/",
	  "timestamp": "2016-07-20T05:54:02.258074Z",
	  "channelId": "emulator",
	  "from": {
		"id": "5e4f5dfa",
		"name": "Bot1"
	  },
	  "conversation": {
		"id": "1cf91be5",
		"name": "Conv1",
		"isGroup": false
	  },
	  "recipient": {
		"id": "617d3bf8",
		"name": "User1"
	  },
	  "locale": "",
	  "text": "Hello User!",
	  "summary": "",
	  "textFormat": "",
	  "attachmentLayout": "",
	  "attachments": [

	  ],
	  "entities": [

	  ],
	  "channelData": null,
	  "replyToId": "",
	  "mentions": null
	}`

	conversationParams = ConversationParameters{
		Bot: ChannelAccount{
			Id: "5e4f5dfa",
			Name: "Bot1",
		},
		IsGroup: false,
		Members: []ChannelAccount{
			ChannelAccount{
				Id: "5e4f5dfa",
				Name: "Bot1",
			},
		},
		TopicName: "Default",
	}

	conversationParamsJson = `
	{
	   "bot":{
		  "id":"5e4f5dfa",
		  "name":"Bot1"
	   },
	   "isGroup":false,
	   "members":[
		  {
			 "id":"5e4f5dfa",
			 "name":"Bot1"
		  }
	   ],
	   "topicName":"Default"
	}`

	attachmentUpload = AttachmentUpload{
		Type: "picture",
		Name: "picture",
		OriginalBase64: []byte("hello"),
		ThumbnailBase64: []byte("good morning"),
	}

	attachmentUploadJson = `
	{
		"type": "picture",
		"name": "picture",
		"originalBase64":"aGVsbG8=",
		"thumbnailBase64":"Z29vZCBtb3JuaW5n"
	}`
)