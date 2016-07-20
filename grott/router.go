package grott

import (
	"fmt"
	"github.com/michael-golfi/Grott/grott/dialog"
	"github.com/michael-golfi/Grott/grott/activity"
	"io/ioutil"
	"net/http"
	"github.com/michael-golfi/Grott/alfredo/controller"
)

func ListenAndServe(cont controller.Controller, router *dialog.DialogRouter) error {

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		var message activity.Activity
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error Reading Body: %s", err.Error())
		}



		if err := message.UnmarshalJSON(b); err != nil {
			fmt.Errorf("Error Decoding Json: %s", err.Error())
		}

		cont.Post(&message)
	})

	fmt.Println("Now serving on port: 3978")
	return http.ListenAndServe(":3798", nil)
}
