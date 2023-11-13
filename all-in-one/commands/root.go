package commands

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/pkg/util"
	"github.com/spf13/cobra"
)

// RootCommand returns root CLI application command interface
func Root() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "all-in-one",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(
		versionCommand(),
		startCommand(),
	)

	return cmd
}

func printHeader() {
	util.PrintInfo("all-in-one")
}
