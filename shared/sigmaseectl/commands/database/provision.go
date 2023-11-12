package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/commands/database/common"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/configuration"
	"github.com/spf13/cobra"
)

type provisionDatabaseOptions struct {
	name string
}

func ProvisionCommand(databaseOptions common.DatabaseOptions) *cobra.Command {
	options := provisionDatabaseOptions{}
	_, sugarLogger := logger.CreateProductionLogger()
	defer func() {
		_ = sugarLogger.Sync()
	}()

	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision Database",
		Long:  "Provision Database",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc()

			configurationHelper, err := appsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			if len(databaseOptions.ConnectionString) != 0 {
				config.Postgres.ConnectionString = databaseOptions.ConnectionString
			}

			database, err := appsetup.NewDatabase(
				sugarLogger,
				config.Postgres,
				config.App)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer database.Close()

			var found int

			db := database.GetDB()
			if err := db.QueryRowContext(
				ctx,
				fmt.Sprintf(
					"SELECT 1 FROM pg_database WHERE datname = '%s'",
					options.name)).
				Scan(&found); err == sql.ErrNoRows {
				if _, err = db.ExecContext(
					ctx,
					fmt.Sprintf("CREATE DATABASE \"%s\"", options.name)); err != nil {
					sugarLogger.Fatal(err)
				}

				sugarLogger.Infof("Database %s successfully provisioned.", options.name)
			} else if err != nil {
				sugarLogger.Fatal(err)
			} else {
				sugarLogger.Infof("Database %s already exists. Ignore provisioning the database.", options.name)
			}
		},
	}

	cmd.Flags().StringVar(&options.name, "name", "", "Specify the database name to provision")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
