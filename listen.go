package bot

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func setup() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/root")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err.Error())
	}
}

type Bot interface {
	Initialize(connector *Connector)

	SetToken(token string)
	GetSession() Session

	Post(session *Session, activity *Activity)
}

func Listen(bot Bot) {
	setup()

	name := viper.GetString("app.name")
	port := viper.GetString("app.port")

	path := "/api/messages"

	conn := &Connector{}

	router := mux.NewRouter()
	router.
		Handle(path, activity(bot, conn)).
		Methods("POST")

	log.Printf("Started %s at %s", name, port)
	log.Fatal(http.ListenAndServe(port, router))
}
