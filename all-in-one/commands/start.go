package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"

	customerapicommands "github.com/sigmasee/sigmasee/customer/customer-api/commands"
	customerapiconfiguration "github.com/sigmasee/sigmasee/customer/customer-api/configuration"
	customerprocessorcommands "github.com/sigmasee/sigmasee/customer/customer-processor/commands"
	customerprocessorconfiguration "github.com/sigmasee/sigmasee/customer/customer-processor/configuration"
	customerctlkafkacommands "github.com/sigmasee/sigmasee/customer/customerctl/commands/kafka"
	customerctlconfiguration "github.com/sigmasee/sigmasee/customer/customerctl/configuration"

	apexapicommands "github.com/sigmasee/sigmasee/apex/apex-api/commands"
	apexapiconfiguration "github.com/sigmasee/sigmasee/apex/apex-api/configuration"

	gatewayapicommands "github.com/sigmasee/sigmasee/gateway/gateway-api/commands"
	gatewayapiconfiguration "github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start all-in-one",
		Long:  "Start all-in-one",
		Run: func(cmd *cobra.Command, args []string) {
			logger, sugarLogger := logger.CreateDevelopmentLogger()
			defer func() {
				_ = sugarLogger.Sync()
			}()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
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

				sugarLogger.Info("Stopping...")

				cancelFunc()
			}()

			/*********************  create topics ****************************/

			/*****************************************************************/
			var customerctlConfig customerctlconfiguration.Config
			if err := configurationHelper.LoadYaml("../customer/customerctl/config.yaml", &customerctlConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			customerctlkafkacommands.CreateTopics(ctx, sugarLogger, customerctlConfig)
			/*****************************************************************/

			/*********************  start components *************************/
			/*****************************************************************/
			var customerApiConfig customerapiconfiguration.Config
			if err := configurationHelper.LoadYaml("../customer/customer-api/config.yaml", &customerApiConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			go customerapicommands.Start(ctx, sugarLogger, customerApiConfig)
			/*****************************************************************/

			/*****************************************************************/
			var customerProcessorConfig customerprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("../customer/customer-processor/config.yaml", &customerProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			go customerprocessorcommands.Start(ctx, sugarLogger, customerProcessorConfig)
			/*****************************************************************/

			/*****************************************************************/
			var apexApiConfig apexapiconfiguration.Config
			if err := configurationHelper.LoadYaml("../apex/apex-api/config.yaml", &apexApiConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			go apexapicommands.Start(ctx, sugarLogger, apexApiConfig)
			/*****************************************************************/

			/*****************************************************************/
			var gatewayApiConfig gatewayapiconfiguration.Config
			if err := configurationHelper.LoadYaml("../gateway/gateway-api/config.yaml", &gatewayApiConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			go gatewayapicommands.Start(ctx, logger, sugarLogger, gatewayApiConfig)
			/*****************************************************************/

			contextHelper, err := enterpriseappsetup.NewContextHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			contextHelper.WaitUntilCancelled(ctx)

			sugarLogger.Info("Stopped")
		},
	}

	return cmd
}
