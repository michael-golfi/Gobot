package grott

import (
	"encoding/json"
	"fmt"
	"github.com/michael-golfi/Grott/grott/types"
	"net/http"
	"github.com/michael-golfi/Grott/grott/dialog"
)

func ListenAndServe(cont types.Controller, router *dialog.DialogRouter) error {

	//controller.StartBot(dialog)

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		var message types.Message

		if err := json.NewDecoder(r.Body).Decode(message); err != nil {
			fmt.Errorf("Error Decoding Json: %s", err.Error())
		}

		response := cont.Post(&message)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Errorf("Error Encoding Json: %s", err.Error())
		}
	})

	return http.ListenAndServe(":9000", nil)
}
