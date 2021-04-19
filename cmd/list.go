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

// listCmd ::: Print the Rainbow
// TODO: ouptut format flag
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
		list, rerr := listRainbow(viper.AllSettings())
		if rerr != nil {
			log.Fatalf("Error on Return: %s", rerr)
		}
		fmt.Println(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// listRainbow ::: Display the entire Rainbow dataset as TOML.
func listRainbow(as map[string]interface{}) (string, error) {
	r, err := toml.Marshal(as)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return string(r), err
	}
	return string(r), nil
}
