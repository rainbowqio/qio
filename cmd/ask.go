/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask QIO a question.",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("ask called")
		readRainbow()
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

// So far, this will only read a single dot. 'meta.editor' will work in the yaml,
// but 'meta.last.editor' will throw an error: While parsing config: (11, 1): unexpected token
// In other words, tables with keys can be [meta.last]:editor but not [meta]:last.editor
// This may actually be more readable.
//
// So the job then is to be able to search the config, which probably means... ???
func readRainbow() {
	if viper.IsSet("blameless.gcp.editor") {
		editor := viper.Get("blameless.gcp.editor")
		fmt.Println(editor)
	} else {
		editor := "ENOENT"
		fmt.Println(editor)
	}
}
