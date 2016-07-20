package dialog

/*

	TODO - Need to reimplement LUIS Dialog Foundation for Bot Framework V3

import (
	"github.com/michael-golfi/Grott/grott/types"
	"errors"
	"net/url"
	"log"
	"net/http"
	"encoding/json"
)

const LUIS_API_BASE string = "https://api.projectoxford.ai/luis/v1/application?"

type LuisDialog struct {
	AppId          string
	SubscriptionId string
	Functions      map[string]func(*types.DialogContext, *types.LuisResult) (*types.Message, error)
}

func (d LuisDialog) HandleFunc(name string, f func(*types.DialogContext, *types.LuisResult) (*types.Message, error)) error {
	if d.Functions[name] != nil {
		return errors.New("Function already exists")
	}
	d.Functions[name] = f
	return nil
}

func (d LuisDialog) MessageReceived(ctx *types.DialogContext, msg types.Message) (*types.Message, error) {
	parameters := map[string]string{
		"id": d.AppId,
		"subscription-key": d.SubscriptionId,
		"q": msg.Text,
	}

	url := buildQueryString(LUIS_API_BASE, parameters)
	resp, err := http.Get(url.RequestURI())

	if err != nil {
		return nil, errors.New("Could not reach Luis Service")
	}

	var luisResult *types.LuisResult
	if err := json.NewDecoder(resp.Body).Decode(luisResult); err != nil {
		return nil, errors.New("Could not decode LuisResult")
	}

	funcName := getLuisIntent(luisResult)
	luisFunc := d.Functions[funcName]
	return luisFunc(ctx, luisResult)
}

func (d LuisDialog) CalculateScore(msg types.Message) (int, error) {
	return 100, nil
}

func buildQueryString(base string, parameters map[string]string) *url.URL {
	baseUrl, err := url.Parse(base)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	for k, v := range parameters {
		params.Add(k, v)
	}

	baseUrl.RawQuery = params.Encode()
	return baseUrl
}

func getLuisIntent(luisResult *types.LuisResult) string {
	highest := 0.0
	funcName := ""
	for _, intent := range luisResult.Intents {
		if intent.Score > highest {
			highest = intent.Score
			funcName = intent.Intent
		}
	}
	return funcName
}*/
