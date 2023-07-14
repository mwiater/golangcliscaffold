package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Path string
var Filecount int
var Minfilesize int64

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Show the largest files in the given path.",
	Long:  `Quickly scan a directory and find large files. . Use the flags below to target the output.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Debug {
			for key, value := range viper.GetViper().AllSettings() {
				log.WithFields(log.Fields{
					key: value,
				}).Info("Command Flag")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)

	filesCmd.PersistentFlags().IntVarP(&Filecount, "filecount", "f", 10, "Limit the number of files returned")
	viper.BindPFlag("filecount", filesCmd.PersistentFlags().Lookup("filecount"))

	filesCmd.PersistentFlags().Int64VarP(&Minfilesize, "minfilesize", "", 50, "Minimum size for files in search in MB.")
	viper.BindPFlag("minfilesize", filesCmd.PersistentFlags().Lookup("minfilesize"))
}
