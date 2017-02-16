package main

import (
	"encoding/json"

	"github.com/mochacoder/goserverless/gsl"
	"github.com/mochacoder/goserverless/gsl/logs"
)

type input struct {
	Word string `json:"word"`
}

type Output struct {
	Length int `json:"length"`
}

func main() {
	goserverless.Handle(func(event json.RawMessage, logger logs.Logger) (interface{}, error) {
		var i input
		var output Output

		err := json.Unmarshal(event, &i)
		if err != nil {
			return nil, err
		}

		output.Length = len(i.Word)

		return output, nil
	})
}
