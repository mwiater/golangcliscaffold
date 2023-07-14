package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Show the largest files in the given path.",
	Long:  `Quickly scan a directory and find large files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("files called")
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
