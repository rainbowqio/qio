/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"errors"
	"fmt"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask QIO a question with: <almanac> <plug>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
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
		switch len(args) {
		case 1:
			readRainbow(args[0], args[0])
		default:
			readRainbow(args[0], args[1])
		}
		/*
			if len(args) == 1 {
				readRainbow(args[0], args[0])
			} else {
				readRainbow(args[0], args[1])
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

// readRainbow ::: Get the value of `almanac.plug`.
func readRainbow(almanac string, plug string) {
	var acPlug string

	if almanac == plug {
		acPlug = almanac
	} else {
		acPlug = almanac + "." + plug
	}

	if viper.IsSet(acPlug) {
		found := viper.Get(acPlug)
		fmtfd, err := toml.Marshal(found)
		// Note that 'dt = 2021-03-30T22:00:00Z' does not throw this error
		// TODO: make this more data-aware?
		if err != nil {
			// The result cannot be marshalled, it isn't TOML so use the original result.
			fmt.Printf("%s ::: %s\n", acPlug, found)
		} else {
			fmt.Printf("%s >>>\n%s\n", acPlug, fmtfd)
		}
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
