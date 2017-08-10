package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func activity(bot Bot) http.Handler {
	mutex := sync.Mutex{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := getSession(bot, &mutex)

		var dialog Activity
		if err := json.NewDecoder(r.Body).Decode(&dialog); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("Bad Request: %s", err.Error())
			return
		}

		bot.Post(session, &dialog)
	})
}

func getSession(bot Bot, mutex *sync.Mutex) *Session {
	mutex.Lock()
	defer mutex.Unlock()

	session := bot.GetSession()
	return &session
}
