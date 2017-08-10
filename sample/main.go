package main

import (
	"log"

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

		break

	default:
		log.Printf("Unsupported Message Type: %s\n", activity.Type)
	}
}

func main() {
	sampleBot := SampleBot{}
	bot.Listen(sampleBot)
}
