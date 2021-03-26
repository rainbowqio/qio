package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number.",
	Long:  `The QIO command lineage in digits.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rainbow QIO Reader v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
