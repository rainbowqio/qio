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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		// When [tab] [tab] is typed for completion, return a list of all Almanac topics in the Rainbow.
		return listAlmanac(), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Args: Almanac, Plug
		readRainbow(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

// readRainbow ::: Get the value of `almanac.plug`.
func readRainbow(almanac string, plug string) {
	acPlug := almanac + "." + plug
	if viper.IsSet(acPlug) {
		found := viper.Get(acPlug)
		fmt.Printf("%s ::: %s\n", acPlug, found)
	} else {
		fmt.Println("ENOENT")
	}
}

// listAlmanac ::: Display the top-level Almanac topics in the Rainbow.
func listAlmanac() []string {
	vas := viper.AllSettings()
	var almanac_list []string
	for key, _ := range vas {
		almanac_list = append(almanac_list, key)
	}
	return almanac_list
}
