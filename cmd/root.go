package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool
var Debug bool
var Highlight int

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
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().IntVarP(&Highlight, "highlight", "", 500, "Highlight files/directories over this threshold, in MB")
	viper.BindPFlag("highlight", rootCmd.PersistentFlags().Lookup("highlight"))

	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", os.Getenv("HOME"), "Define the path to scan.")
	rootCmd.MarkFlagRequired("path")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))

}
