package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sigmasee/sigmasee/customer/customer-processor/appsetup"
	"github.com/sigmasee/sigmasee/customer/customer-processor/configuration"
	customerappsetup "github.com/sigmasee/sigmasee/customer/shared/appsetup"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start customer-processor",
		Long:  "Start customer-processor",
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

	customerConsumer, err := appsetup.NewCustomerConsumer(
		sugarLogger,
		config.App,
		config.Kafka,
		entgoClient)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if err = customerConsumer.StartAsync(ctx); err != nil {
		sugarLogger.Fatal(err)
	}

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if err = httpServer.ListenAndServe(); err != nil {
		sugarLogger.Fatal(err)
	}
}
