package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/michael-golfi/Grott"
	"github.com/michael-golfi/Grott/storage"
)

type SampleBot struct {
}

func (b SampleBot) Initialize() {

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

			bot.CreateConversation(activity, &conversationParameters)
		case "get_activity_members":
			m, err := bot.GetActivityMembers(activity)
			spew.Dump(m, err)
			break

		case "get_conversation_members":
			m, err := bot.GetConversationMembers(activity)
			spew.Dump(m, err)
			break

		case "send":
			mockActivity := activity.CreateReply("Hello User!", "en-US")
			bot.SendToConversation(&mockActivity)
			break

		case "reply":
			mockActivity := activity.CreateReply("Hello User!", "en-US")
			bot.ReplyToActivity(&mockActivity)
			break

		case "update":
			break

		case "delete":
			bot.DeleteActivity(activity)
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
