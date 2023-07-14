package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool
var Debug bool

var rootCmd = &cobra.Command{
	Use:   "getsize",
	Short: "List the size of a local directory.",
	Long:  `This command will display the size of a directory with several different options.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
