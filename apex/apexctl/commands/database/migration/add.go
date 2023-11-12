package migration

import (
	"context"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/sigmasee/sigmasee/apex/apexctl/configuration"
	"github.com/sigmasee/sigmasee/apex/shared/appsetup"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
)

type addMigrationOptions struct {
	path                 string
	name                 string
	enableGlobalUniqueID bool
	enableDropColumn     bool
	enableDropIndex      bool
}

func AddCommand(connectionString *string) *cobra.Command {
	options := addMigrationOptions{}

	_, sugarLogger := logger.CreateProductionLogger()
	defer func() {
		_ = sugarLogger.Sync()
	}()

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add new migration",
		Long:  "Add new migration",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			osHelper, err := enterpriseappsetup.NewOsHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			if len(*connectionString) != 0 {
				config.Postgres.ConnectionString = *connectionString
			}

			entgoClient, err := appsetup.NewEntgoClient(
				sugarLogger,
				config.Database,
				config.Postgres,
				config.App)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer entgoClient.Close()

			if !osHelper.DirExist(options.path) {
				if err := osHelper.CreateDir(options.path); err != nil {
					sugarLogger.Fatalf("failed to create migration directory %s. Error: %v", options.path, err)
				}
			}

			migrationDirectory, err := atlas.NewLocalDir(options.path)
			if err != nil {
				sugarLogger.Fatalf("failed creating atlas migration directory %s. Error: %v", options.path, err)
			}

			migrationOptions := []schema.MigrateOption{
				schema.WithDir(migrationDirectory),
				schema.WithMigrationMode(schema.ModeInspect),
				schema.WithDialect(dialect.Postgres),
				schema.WithFormatter(atlas.DefaultFormatter),
				schema.WithGlobalUniqueID(options.enableGlobalUniqueID),
				schema.WithDropColumn(options.enableDropColumn),
				schema.WithDropIndex(options.enableDropIndex),
			}

			client := entgoClient.GetClient()

			if err = client.Schema.NamedDiff(
				ctx,
				options.name,
				migrationOptions...); err != nil {
				sugarLogger.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&options.path, "path", "", "Specify the path to the migration scripts directory")
	cmd.Flags().StringVar(&options.name, "name", "", "Specify the name of the migration to add")
	cmd.Flags().BoolVar(&options.enableGlobalUniqueID, "enable-global-unique-id", false, "Adds the universal ids options to the migration.")
	cmd.Flags().BoolVar(&options.enableDropColumn, "enable-drop-column", true, "Adds the column dropping options to the migration.")
	cmd.Flags().BoolVar(&options.enableDropIndex, "enable-drop-index", true, "Adds the index dropping options to the migration.")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("name"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
