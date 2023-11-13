package commands

import (
	"github.com/sigmasee/sigmasee/customer/customerctl/commands/kafka"
	"github.com/sigmasee/sigmasee/customer/customerctl/commands/kafka/common"
	"github.com/spf13/cobra"
)

func kafkaCommand() *cobra.Command {
	options := common.KafkaOptions{}

	cmd := &cobra.Command{
		Use: "kafka",
	}

	cmd.PersistentFlags().StringVar(&options.BootstrapServers, "bootstrapServers", "", "Specify the Kafka bootstrap servers")

	cmd.AddCommand(
		kafka.CreateTopicsCommand(options),
	)

	return cmd
}
