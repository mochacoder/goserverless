package main

import (
	"encoding/json"

	"github.com/mochacoder/goserverless/goserverless-go"
	"github.com/mochacoder/goserverless/goserverless-go/logs"
)

func main() {
	goserverless.Handle(func(raw json.RawMessage, logger logs.Logger) (interface{}, error) {
		return "Hello World!", nil
	})
}
