package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/michael-golfi/Grott"
	"github.com/michael-golfi/Grott/storage"
)

type SampleBot struct {
	Connector *bot.Connector
	token     string
}

func (b SampleBot) Initialize(connector *bot.Connector) {
	b.Connector = connector
}

func (b SampleBot) SetToken(token string) {
	b.token = token
}

func (b SampleBot) GetSession() bot.Session {
	return storage.SessionMemory{}
}

func (b SampleBot) Post(session *bot.Session, activity *bot.Activity) {
	switch activity.Type {
	case "conversationUpdate":
		break

	case "message":

		switch activity.Text {

		case "create":
			conversationParameters := bot.ConversationParameters{
				Bot:       activity.Recipient,
				IsGroup:   activity.Conversation.IsGroup,
				Members:   []bot.ChannelAccount{activity.From},
				TopicName: activity.TopicName,
			}

			b.Connector.CreateConversation(activity, &conversationParameters)
		case "get_activity_members":
			m, err := b.Connector.GetActivityMembers(activity)
			spew.Dump(m, err)
			break

		case "get_conversation_members":
			m, err := b.Connector.GetConversationMembers(activity)
			spew.Dump(m, err)
			break

		case "send":
			mockActivity := activity.CreateReply("Hello User!", "en-US")
			b.Connector.SendToConversation(&mockActivity)
			//spew.Dump(activity, mockActivity)
			break

		case "reply":
			mockActivity := activity.CreateReply("Hello User!", "en-US")
			b.Connector.ReplyToActivity(&mockActivity)

			//spew.Dump(activity, mockActivity)
			break

		case "update":
			break

		case "delete":
			b.Connector.DeleteActivity(activity)
			break

		default:
			break

		}

		break

	default:
		log.Printf("Unsupported Message Type: %s\n", activity.Type)
	}
}

func main() {
	sampleBot := SampleBot{}
	bot.Listen(sampleBot)
}
