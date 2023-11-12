package generate

import (
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands/generate/client"
	"github.com/spf13/cobra"
)

func ClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "client",
	}

	cmd.AddCommand(
		client.EventCommand(),
	)

	return cmd
}
