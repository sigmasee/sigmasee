package kafka

import (
	"context"

	customerctlappsetup "github.com/sigmasee/sigmasee/customer/customerctl/appsetup"
	"github.com/sigmasee/sigmasee/customer/customerctl/commands/kafka/common"
	"github.com/sigmasee/sigmasee/customer/customerctl/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func CreateTopicsCommand(kafkaOptions common.KafkaOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-topics",
		Short: "Create topics",
		Long:  "Create topics",
		Run: func(cmd *cobra.Command, args []string) {
			_, sugarLogger := logger.CreateProductionLogger()
			defer func() {
				_ = sugarLogger.Sync()
			}()

			configurationHelper, err := appsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			if len(kafkaOptions.BootstrapServers) != 0 {
				config.Kafka.BootstrapServers = kafkaOptions.BootstrapServers
			}

			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc()

			CreateTopics(ctx, sugarLogger, config)
		},
	}

	return cmd
}

func CreateTopics(ctx context.Context, sugarLogger *zap.SugaredLogger, config configuration.Config) {
	if topicService, err := customerctlappsetup.NewTopicService(sugarLogger, config.Kafka); err != nil {
		sugarLogger.Fatal(err)
	} else {
		if err := topicService.CreateTopics(ctx); err != nil {
			sugarLogger.Fatal(err)
		}
	}
}
