/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var format string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export QIO knowledge",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Command does not take an argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigType(format)
		viper.SafeWriteConfigAs("rainbow." + format)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&format, "format", "f", "toml", "export format")
}
