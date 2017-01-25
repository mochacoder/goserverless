package main

import (
	"encoding/json"

	"github.com/mochacoder/goserverless/goserverless-go"
	"github.com/mochacoder/goserverless/goserverless-go/logs"
)

type event struct {
	Name string `json:"name"`
}

type response struct {
	Body string `json:"body"`
}

func main() {
	//	goserverless.FunctionName = "lol"
	goserverless.Handle(func(raw json.RawMessage, logger logs.Logger) (interface{}, error) {
		var input event
		err := json.Unmarshal(raw, &input)
		if err != nil {
			return nil, err
		}

		return response{Body: "hello " + input.Name}, nil
	})
}
