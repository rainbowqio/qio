/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
// So far, this is a basic list. It can help debug too.
// TODO: Allow for list args, e.g. `qio list almanacs`
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List what QIO knows",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Command does not take an argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(listRainbow())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// listRainbow ::: Display the entire Rainbow dataset as TOML.
func listRainbow() string {
	v := viper.AllSettings()
	r, err := toml.Marshal(v)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
	return string(r)
}
