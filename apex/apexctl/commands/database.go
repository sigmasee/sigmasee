package commands

import (
	"github.com/sigmasee/sigmasee/apex/apexctl/commands/database"
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
	)

	return cmd
}
