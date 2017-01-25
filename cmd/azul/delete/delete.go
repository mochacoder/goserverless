package deploy

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/mochacoder/goserverless/cmd/gsl/root"
	"github.com/mochacoder/goserverless/cmd/gsl/utils"
)

var cmd = &cobra.Command{
	Use:     "delete <funcName>",
	Short:   "deprovision the specified function",
	PreRunE: preRun,
	RunE:    run,
}

func init() {
	root.RootCmd.AddCommand(cmd)
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

	return app.Delete(args[0])
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
		return errors.New("delete should be call with one argument")
	}

	return nil
}
