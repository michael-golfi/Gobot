package grott

import (
	"encoding/json"
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
	"net/http"
	"github.com/michael-golfi/Grott/grott/dialog"
	"io/ioutil"
)

func ListenAndServe(cont types.Controller, router *dialog.DialogRouter) error {

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		var message types.Message
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error Reading Body: %s", err.Error())
		}

		if err := json.Unmarshal(b, &message); err != nil {
			fmt.Errorf("Error Decoding Json: %s", err.Error())
		}

		response, err := cont.Post(&message)

		if err != nil {
			// TODO - Good idea to return err to chat?
			response = &types.Message{
				Text: err.Error(),
			}
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Errorf("Error Encoding Json: %s", err.Error())
		}
	})

	fmt.Println("Now serving on port: 3978")
	return http.ListenAndServe(":3978", nil)
}