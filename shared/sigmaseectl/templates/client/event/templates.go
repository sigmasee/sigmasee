package clienteventtemplates

import (
	_ "embed"
)

//go:embed metadata.tmpl
var metadata string

func GetMetadata() string {
	return metadata
}

//go:embed handler.tmpl
var handler string

func GetHandler() string {
	return handler
}

//go:embed consumer.tmpl
var consumer string

func GetConsumer() string {
	return consumer
}

//go:embed consumer_aws_lambda.tmpl
var consumerAwsLambda string

func GetConsumerAwsLambda() string {
	return consumerAwsLambda
}

//go:embed producer.tmpl
var producer string

func GetProducer() string {
	return producer
}

//go:embed generate.tmpl
var generate string

func GetGenerate() string {
	return generate
}
