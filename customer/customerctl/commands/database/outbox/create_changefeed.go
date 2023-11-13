package outbox

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/sigmasee/sigmasee/customer/customerctl/configuration"
	customeroutboxentity "github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
)

type createChangeFeedOptions struct {
	webhookUrl string
}

func CreateChangeFeedCommand(connectionString *string) *cobra.Command {
	options := createChangeFeedOptions{}

	_, sugarLogger := logger.CreateProductionLogger()
	defer func() {
		_ = sugarLogger.Sync()
	}()

	cmd := &cobra.Command{
		Use:   "create-changefeed",
		Short: "Create changefeed for outbox table",
		Long:  "Create changefeed for outbox table",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
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

			database, err := enterpriseappsetup.NewDatabase(
				sugarLogger,
				config.Postgres,
				config.App)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer database.Close()

			db := database.GetDB()

			webhookUrl, err := url.JoinPath(options.webhookUrl, "customer/api/v1/outbox-cdc")
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var jobID int64

			if err := db.QueryRowContext(
				ctx,
				fmt.Sprintf(
					"WITH changefeeds as (SHOW CHANGEFEED JOBS) SELECT job_id FROM changefeeds WHERE sink_uri = ('%s') AND status = ('running');",
					webhookUrl)).
				Scan(&jobID); err == sql.ErrNoRows {

				sugarLogger.Info("No changefeed job found, creating one...")

				if _, err = db.ExecContext(
					ctx,
					fmt.Sprintf("CREATE CHANGEFEED FOR TABLE %s INTO '%s'",
						customeroutboxentity.Table,
						webhookUrl)); err != nil {
					sugarLogger.Fatal(err)
				}

				sugarLogger.Info("Successfully created changefeed job.")
			} else if err != nil {
				sugarLogger.Fatal(err)
			} else {
				sugarLogger.Infof("An existing running job already exists with id: %d, no need to create a new changefeed job", jobID)
			}
		},
	}

	cmd.Flags().StringVar(&options.webhookUrl, "webhook_url", "", "Specify the webhook URL to call when a change happens in outbox table")

	if err := cmd.MarkFlagRequired("webhook_url"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
