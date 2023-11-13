package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sigmasee/sigmasee/customer/customer-api/appsetup"
	"github.com/sigmasee/sigmasee/customer/customer-api/configuration"
	customerappsetup "github.com/sigmasee/sigmasee/customer/shared/appsetup"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start customer-api",
		Long:  "Start customer-api",
		Run: func(cmd *cobra.Command, args []string) {
			_, sugarLogger := logger.CreateProductionLogger()
			defer func() {
				_ = sugarLogger.Sync()
			}()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc()

			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			go func() {
				<-sigc
				cancelFunc()
			}()

			Start(ctx, sugarLogger, config)
		},
	}

	return cmd
}

func Start(ctx context.Context, sugarLogger *zap.SugaredLogger, config configuration.Config) {
	entgoClient, err := customerappsetup.NewEntgoClient(
		sugarLogger,
		config.Database,
		config.Postgres,
		config.App)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	jobScheduler := gocron.NewScheduler(time.UTC)
	jobScheduler.StartAsync()
	defer jobScheduler.Stop()

	tokenService, err := enterpriseappsetup.NewTokenService(
		ctx,
		sugarLogger,
		config.CognitoIdentityProvider,
		config.GoogleIdentityProvider,
		config.SlackIdentityProvider,
		jobScheduler)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	kafkaClient, err := enterpriseappsetup.NewKafkaGoKafkaClient(
		sugarLogger,
		config.Kafka)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if _, err = customerappsetup.NewOutboxBackgroundService(
		ctx,
		sugarLogger,
		kafkaClient,
		config.Outbox,
		entgoClient,
		jobScheduler); err != nil {
		sugarLogger.Fatal(err)
	}

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App,
		config.Outbox,
		config.Intercom,
		entgoClient,
		tokenService,
		kafkaClient)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if err = httpServer.ListenAndServe(); err != nil {
		sugarLogger.Fatal(err)
	}
}
