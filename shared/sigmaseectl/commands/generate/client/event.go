package client

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/appsetup"
	"github.com/spf13/cobra"
)

type eventOptions struct {
	protobufFilePath     string
	packageName          string
	eventType            string
	topicName            string
	retryTopicNamePrefix string
	retryTopicNameCount  int
	deadLetterTopicName  string
	outputPath           string
}

func EventCommand() *cobra.Command {
	options := eventOptions{}
	_, sugarLogger := logger.CreateProductionLogger()
	defer func() {
		_ = sugarLogger.Sync()
	}()

	cmd := &cobra.Command{
		Use:   "event",
		Short: "Event",
		Long:  "Event",
		Run: func(cmd *cobra.Command, args []string) {
			clientEventSchemaGeneratorService, err := appsetup.NewClientEventSchemaGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventMetadataGeneratorService, err := appsetup.NewClientEventMetadataGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventHandlerGeneratorService, err := appsetup.NewClientEventHandlerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventConsumerGeneratorService, err := appsetup.NewClientEventConsumerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventConsumerAwsLambdaGeneratorService, err := appsetup.NewClientEventConsumerAwsLambdaGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventProducerGeneratorService, err := appsetup.NewClientEventProducerGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			clientEventGenerateGeneratorService, err := appsetup.NewClientEventGenerateGeneratorService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventSchemaGeneratorService.Generate(
				options.protobufFilePath,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventMetadataGeneratorService.Generate(
				options.packageName,
				options.topicName,
				options.retryTopicNamePrefix,
				options.retryTopicNameCount,
				options.deadLetterTopicName,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventHandlerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventConsumerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventConsumerAwsLambdaGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventProducerGeneratorService.Generate(
				options.packageName,
				options.eventType,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = clientEventGenerateGeneratorService.Generate(
				options.packageName,
				options.outputPath); err != nil {
				sugarLogger.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&options.protobufFilePath, "protobufFilePath", "", "Specify the protobuf file path")
	cmd.Flags().StringVar(&options.packageName, "packageName", "", "Specify the package name")
	cmd.Flags().StringVar(&options.eventType, "eventType", "", "Specify the event type")
	cmd.Flags().StringVar(&options.topicName, "topicName", "", "Specify the topic name")
	cmd.Flags().StringVar(&options.retryTopicNamePrefix, "retryTopicNamePrefix", "", "Specify the retry topic name")
	cmd.Flags().IntVar(&options.retryTopicNameCount, "retryTopicNameCount", 1, "Specify the retry topic count")
	cmd.Flags().StringVar(&options.deadLetterTopicName, "deadLetterTopicName", "", "Specify the dead letter topic name")
	cmd.Flags().StringVar(&options.outputPath, "outputPath", "", "Specify the output path")

	if err := cmd.MarkFlagRequired("protobufFilePath"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("packageName"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("eventType"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("topicName"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("retryTopicNamePrefix"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("deadLetterTopicName"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("retryTopicNameCount"); err != nil {
		sugarLogger.Fatal(err)
	}

	if err := cmd.MarkFlagRequired("outputPath"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
