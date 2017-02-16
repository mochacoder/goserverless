package main

import (
	"encoding/json"

	"github.com/mochacoder/goserverless/gsl"
	"github.com/mochacoder/goserverless/gsl/logs"
)

func main() {
	goserverless.Handle(func(raw json.RawMessage, logger logs.Logger) (interface{}, error) {
		return "Hello World!", nil
	})
}
