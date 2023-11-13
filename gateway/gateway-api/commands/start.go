package commands

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sigmasee/sigmasee/gateway/gateway-api/appsetup"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start gateway-api",
		Long:  "Start gateway-api",
		Run: func(cmd *cobra.Command, args []string) {
			logger, sugarLogger := logger.CreateProductionLogger()
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

			Start(ctx, logger, sugarLogger, config)
		},
	}

	return cmd
}

func Start(ctx context.Context, logger *zap.Logger, sugarLogger *zap.SugaredLogger, config configuration.Config) {
	httpServer, err := appsetup.NewHttpServer(
		ctx,
		sugarLogger,
		logger,
		config.App,
		config.ApiEndpoints,
		&http.Client{
			Timeout: config.ApiEndpoints.Timeout,
		})
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if err = httpServer.ListenAndServe(ctx); err != nil {
		sugarLogger.Fatal(err)
	}
}
