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
		listAlmanac()
		readRainbow(args[0], args[1])
	},
	ValidArgs: []string{"craque", "mattic"},
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

type Almanac struct {
	Name string
	Plug string
}

type Rainbow struct {
	Almanacs []Almanac
}

var rainbow *Rainbow

func listAlmanac() *Rainbow {
	vas := viper.AllSettings()
	// how do i extract the keys from this
	almanacs := make([]Almanac, 0, len(vas))
	for key, value := range vas {
		val := value.(map[string]interface{})
		fmt.Printf("%s ::: %s\n", key, val)
		// append and return below
		almanacs = append(almanacs, Almanac{
			Name: key,
			Plug: val["*"].(string),
		})
	}
	rainbow = &Rainbow{almanacs}
	return rainbow
}
