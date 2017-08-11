package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type login struct {
	Type         string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}

var token = login{
	ExpiresIn: 0,
}

func authenticate(id, secret string) (*login, error) {
	body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=https://api.botframework.com/.default", id, secret)
	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token", strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var l login
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		return nil, err
	}

	return &l, nil
}

func getToken() (string, error) {
	if token.ExpiresIn <= 0 {
		appID := viper.GetString("app.appId")
		appSecret := viper.GetString("app.appSecret")
		tok, err := authenticate(appID, appSecret)
		if err != nil {
			return "", nil
		}

		token.AccessToken = tok.AccessToken
		token.ExpiresIn = tok.ExpiresIn
		token.ExtExpiresIn = tok.ExtExpiresIn
		token.Type = tok.Type
	}

	return token.AccessToken, nil
}
