package main

import (
	_ "github.com/mochacoder/goserverless/cmd/gsl/delete"
	_ "github.com/mochacoder/goserverless/cmd/gsl/deploy"
	_ "github.com/mochacoder/goserverless/cmd/gsl/logs"
	"github.com/mochacoder/goserverless/cmd/gsl/root"
)

func main() {
	root.Execute()
}
