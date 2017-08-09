package main

import (
	"fmt"

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

func (b SampleBot) Post(session *bot.Session) {
	// Do something
	fmt.Println("Hello Bot")
}

func main() {
	sampleBot := SampleBot{}
	bot.Listen(sampleBot)
}
