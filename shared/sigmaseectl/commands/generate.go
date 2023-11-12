package commands

import (
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands/generate"
	"github.com/spf13/cobra"
)

func generateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "generate",
	}

	cmd.AddCommand(
		generate.ClientCommand(),
	)

	return cmd
}
