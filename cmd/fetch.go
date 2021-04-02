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

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch what QIO knows",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Usage: qio fetch")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fetchRainbow())
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func fetchRainbow() string {
	v := viper.AllSettings()
	r, err := toml.Marshal(v)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
	return string(r)
}
