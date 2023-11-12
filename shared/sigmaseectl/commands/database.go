package commands

import (
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands/database"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands/database/common"
	"github.com/spf13/cobra"
)

func databaseCommand() *cobra.Command {
	options := common.DatabaseOptions{}

	cmd := &cobra.Command{
		Use: "database",
	}

	cmd.PersistentFlags().StringVar(&options.ConnectionString, "connectionString", "", "Specify the database connection string")

	cmd.AddCommand(
		database.ProvisionCommand(options),
	)

	return cmd
}
