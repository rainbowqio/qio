/*
Copyright © 2021 Matt Davis <maroda@rainbowq.io>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script for QIO",
	Long: `QIO can complete some commands with a list of known Almanacs.
To load completions in your shell of choice:

Zsh:

  Load for each session for your user (only necessary once):

      $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  Then export the script (only necessary once):

      $ qio completion zsh > "${fpath[1]}/_qio"

  Start a new shell for this to take effect.

Bash:

  AdHoc for current session only:

      $ source <(qio completion bash)

  Load for each session (only necessary once):
  Linux:

      $ qio completion bash > /etc/bash_completion.d/qio

  macOS:

  	$ qio completion bash > /usr/local/etc/bash_completion.d/qio

fish:

  # AdHoc for current session only:

  	$ qio completion fish | source

  # Load for each session (only necessary once):

  	$ qio completion fish > ~/.config/fish/completions/qio.fish

PowerShell:

  AdHoc for current session only:

  	PS> qio completion powershell | Out-String | Invoke-Expression

  Load for each session (source this file from your PowerShell profile):

  	PS> qio completion powershell > qio.ps1
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
