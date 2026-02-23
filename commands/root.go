package commands

import (
	"fmt"
	"os"

	"github.com/IronFE/iron.cli/commands/ecr"
	"github.com/spf13/cobra"
)

// the version is set via ldflags
var version = "snapshot"

var rootCmd = &cobra.Command{
	Use:   "iron",
	Short: "IronCLI helps you manage your stuff",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

func ExecuteRootCommand() {

	rootCmd.AddCommand(NewPlanCommand())
	rootCmd.AddCommand(NewDeployCommand())
	rootCmd.AddCommand(NewDestroyCommand())
	rootCmd.AddCommand(NewOutputCommand())
	rootCmd.AddCommand(NewAuthorizeCommand())
	rootCmd.AddCommand(NewSsmSessionCommand())
	rootCmd.AddCommand(ecr.NewEcrCommand())

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceUsage = true
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("{{println .Version}}")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
