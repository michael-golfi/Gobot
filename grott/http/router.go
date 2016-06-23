package grott

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func ListenAndServe() {

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		var message Message
		if err := json.NewDecoder(r.Body).Decode(message); err != nil {
			fmt.Errorf("Error Decoding Json: %s", err.Error())
		}

		response := Post(&message)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Errorf("Error Encoding Json: %s", err.Error())
		}
	})

	http.ListenAndServe(":9000", nil)
}