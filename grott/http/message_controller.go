package grott

var (
	Send chan Message
	Receive chan Message
)

func Post(msg *Message) *Message {


	return handleSystemMessage(msg)
}

func handleSystemMessage(msg *Message) *Message {
	if msg.MessageType == "BotAddedToConversation" {
		return &Message{
			Text: "Hello World",
		}
	}

	if msg.MessageType == "UserAddedToConversation" {
		return &Message{
			Text: "Hello Human!",
		}
	}

	if msg.MessageType == "BotRemovedFromConversation" {
		return &Message{
			Text: "Goodbye Bot!",
		}
	}

	if msg.MessageType == "UserRemovedFromConversation" {
		return &Message{
			Text: "Goodbye Human!",
		}
	}

	if msg.MessageType == "EndOfConversation" {
		return &Message{
			Text: "End of Conversation",
		}
	}
}