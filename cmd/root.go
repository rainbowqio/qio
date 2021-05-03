/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qio",
	Short: "RainbowQ ::: all your knowledgebase are belong to us",
	Example: `
	::: Display a Plug within an Almanac:
	$ qio ask <almanac> <plug>

	::: List what QIO knows:
	$ qio list

	::: QIO has randomizers! Toss a coin for HEADS or TAILS, get a token, or get a URL-encoded base64 token:
	$ qio coin
	$ qio coin -t
	$ qio coin -b

	::: Install shell completion (see 'qio completion --help' for more):
	$ qio completion zsh > "${fpath[1]}/_qio"

	::: Export known Rainbow to a local TOML file:
	$ qio export
	`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// cobra.CheckErr(rootCmd.Execute())
	// testing a non-function err here
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Bool("version", false, "print the version")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/qio/rainbow.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Read ~/.config/qio/ or create if missing
		cfgPath := fmt.Sprintf("%s/.config/qio/", home)
		dir := filepath.Clean(cfgPath)
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatalf("error creating config dir: %s", err)
		}
		_, err = os.OpenFile(dir+"/rainbow.toml", os.O_RDONLY|os.O_CREATE, 0660)
		if err != nil {
			log.Fatalf("error creating config file: %s", err)
		}

		// set config in "~/.config/qio/rainbow.toml"
		viper.AddConfigPath(cfgPath)
		viper.SetConfigName("rainbow")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
