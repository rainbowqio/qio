/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask QIO a question with: <almanac> <plug>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Usage: qio ask <almanac> <plug>")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Args: Almanac, Plug
		readRainbow(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

func readRainbow(almanac string, plug string) {
	acPlug := almanac + "." + plug
	if viper.IsSet(acPlug) {
		found := viper.Get(acPlug)
		fmt.Printf("%s ::: %s\n", acPlug, found)
	} else {
		fmt.Println("ENOENT")
	}
}
