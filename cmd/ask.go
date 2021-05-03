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
		return listAlmanacs(viper.AllSettings()), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		var answer string
		switch len(args) {
		case 1:
			// One Arg is taken as an Almanac search
			_, _, answer = readRainbow(args[0], args[0])
		default:
			// Two args are an Almanac + Plug search
			_, _, answer = readRainbow(args[0], args[1])
		}
		fmt.Println(answer)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

// readRainbow ::: Get the value for the requested Almanac:Plug pair.
//		A Rainbow is all the data.
//		A Plug either contains or points to external sources of data.
//		A query/question for the Rainbow looks at the value of a Plug.
//
//      When the args are equal, those are table names.
//      We are showing a full Almanac.
//
//		When the args differ, the first is the Almanac,
// 		  the second is the Plug, or key.
//		This accepts any combination of compound key names,
// 		  with the notion that it is only an Almanac if
//		  it is a top-level TOML table name.
//
func readRainbow(almanac string, plug string) (string, string, string) {
	var acPlug string

	if almanac == plug {
		acPlug = almanac
	} else {
		acPlug = almanac + "." + plug
	}
	display, val, _ := askAlmanac(acPlug)

	switch display {
	case "Not Found":
		return acPlug, display, "Not Found, does this Almanac exist?"
	case "This Level":
		fmt.Printf("Plugs found in %s:\n\n", acPlug)
		fmt.Println(val)
		return acPlug, display, ""
	case "Plug Is":
		//fmt.Println("Answer Found!")
		return acPlug, display, val
	default:
		return acPlug, display, "default"
	}

}

// askAlmanac ::: Validates and retrieves a Rainbow Almanac:Plug (ACPlug)
// 		No matching Almanac Plug, no data.
//		There are no default ACPlugs.
func askAlmanac(acp string) (string, string, error) {
	// Validate first, then move on
	if !viper.IsSet(acp) {
		return "Not Found", "", errors.New("ENOENT")
	}
	found := viper.Get(acp)

	var rVal string

	// When the TOML library does not return an error,
	// the source is marshalled and displayed as TOML.
	fmtfd, err := toml.Marshal(found)
	if err == nil {
		// fmt.Printf("%s >>>\n%s\n", acp, fmtfd)
		rVal = fmt.Sprint(fmtfd)
		return "This Level", string(fmtfd), nil
	}

	// An unmarshallable source is the value of the plug.
	// fmt.Printf("%s ::: %s\n", acp, found)
	rVal = fmt.Sprint(found)
	return "Plug Is", rVal, nil
}

// listAlmanacs ::: Display the top-level Almanac topics in the Rainbow.
// 		An Almanac is a TOML table name. This list can be used to provide
// 		shell tab completion or queries that require these top-level keys.
func listAlmanacs(as map[string]interface{}) []string {
	var almanac_list []string
	for key, _ := range as {
		almanac_list = append(almanac_list, key)
	}
	return almanac_list
}
