package database

import (
	outbox "github.com/sigmasee/sigmasee/customer/customerctl/commands/database/outbox"
	"github.com/spf13/cobra"
)

func OutboxCommand(connectionString *string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "outbox",
	}

	cmd.AddCommand(
		outbox.CreateChangeFeedCommand(connectionString),
	)

	return cmd
}
