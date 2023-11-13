package commands

import (
	"github.com/sigmasee/sigmasee/customer/customerctl/commands/database"
	"github.com/spf13/cobra"
)

type databaseOptions struct {
	connectionString string
}

func databaseCommand() *cobra.Command {
	options := databaseOptions{}

	cmd := &cobra.Command{
		Use: "database",
	}

	cmd.PersistentFlags().StringVar(&options.connectionString, "connectionString", "", "Specify the database connection string")

	cmd.AddCommand(
		database.MigrateCommand(&options.connectionString),
		database.OutboxCommand(&options.connectionString),
	)

	return cmd
}
