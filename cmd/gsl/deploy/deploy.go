package deploy

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/mochacoder/goserverless/cmd/gsl/root"
	"github.com/mochacoder/goserverless/cmd/gsl/utils"
	"github.com/mochacoder/goserverless/function"
)

var deployCmd = &cobra.Command{
	Use:     "deploy <funcName>",
	Short:   "deploy the specified function to Azure",
	PreRunE: preRun,
	RunE:    run,
}

func init() {
	root.RootCmd.AddCommand(deployCmd)
}

func run(cmd *cobra.Command, args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	app, err := utils.GetApp(wd)
	if err != nil {
		return err
	}

	f := function.Function{Name: args[0], Path: wd + "/" + args[0] + "/"}

	return app.Deploy(&f)
}

func preRun(cmd *cobra.Command, args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := os.Stat(wd + "/goserverless.json"); os.IsNotExist(err) {
		return errors.New("goserverless.json file not found")
	}

	if len(args) != 1 {
		return errors.New("deploy should be call with one argument")
	}

	return checkFunctionSanity(wd, args[0])
}

func checkFunctionSanity(wd string, dirname string) error {
	if _, err := os.Stat(wd + "/" + dirname); os.IsNotExist(err) {
		return errors.New("Function " + dirname + " not found")
	}
	return nil
}
