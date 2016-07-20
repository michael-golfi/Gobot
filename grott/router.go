package grott

import (
	"encoding/json"
	"fmt"
	"github.com/michael-golfi/Grott/grott/dialog"
	"github.com/michael-golfi/Grott/grott/types"
	"io/ioutil"
	"net/http"
)

func ListenAndServe(cont types.Controller, router *dialog.DialogRouter) error {

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		var message types.Activity
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error Reading Body: %s", err.Error())
		}

		if err := json.Unmarshal(b, &message); err != nil {
			fmt.Errorf("Error Decoding Json: %s", err.Error())
		}

		cont.Post(&message)
	})

	fmt.Println("Now serving on port: 3978")
	return http.ListenAndServe(":3798", nil)
}
