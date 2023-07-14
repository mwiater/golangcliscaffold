/*
Copyright © 2023 Matt J. Wiater matt.wiater@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "directorylist",
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
