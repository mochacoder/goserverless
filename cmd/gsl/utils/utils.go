package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mochacoder/goserverless/app"
)

//GetApp returns an App based on goserverless.json
func GetApp(wd string) (*app.App, error) {
	file, err := os.Open(wd + "/goserverless.json")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var app app.App
	err = json.Unmarshal(b, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}
