package logs

import (
	"errors"
	"os"

	"github.com/mochacoder/goserverless/cmd/gsl/root"
	"github.com/mochacoder/goserverless/cmd/gsl/utils"
	"github.com/mochacoder/goserverless/function"
	"github.com/spf13/cobra"
)

var follow bool

var logsCmd = &cobra.Command{
	Use:     "logs <funcName>",
	Short:   "Outputs function logs",
	RunE:    run,
	PreRunE: preRun,
}

func init() {
	root.RootCmd.AddCommand(logsCmd)

	f := logsCmd.Flags()
	f.BoolVarP(&follow, "follow", "f", false, "Specify if the logs should be streamed. When using -f, you don't need to specify a function.")
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

	if follow {
		return app.LogStream()
	}

	f := function.Function{Name: args[0], Path: wd + "/" + args[0] + "/"}
	return app.Logs(&f)
}

func preRun(cmd *cobra.Command, args []string) error {
	if (len(args) == 0 && !follow) || (len(args) == 1 && follow) {
		return errors.New("you should either specify the name of the function to log, or add the -f flag")
	}

	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	return nil
}
