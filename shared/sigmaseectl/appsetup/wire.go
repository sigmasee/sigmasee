// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package appsetup

import (
	"github.com/google/wire"
	"github.com/sigmasee/sigmasee/shared/enterprise/os"
	"github.com/sigmasee/sigmasee/shared/sigmaseectl/services/generators"
	"go.uber.org/zap"
)

func NewClientEventSchemaGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventSchemaGeneratorService, error) {
	wire.Build(
		generators.NewClientEventSchemaGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventMetadataGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventMetadataGeneratorService, error) {
	wire.Build(
		generators.NewClientEventMetadataGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventHandlerGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventHandlerGeneratorService, error) {
	wire.Build(
		generators.NewClientEventHandlerGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventConsumerGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventConsumerGeneratorService, error) {
	wire.Build(
		generators.NewClientEventConsumerGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventConsumerAwsLambdaGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventConsumerAwsLambdaGeneratorService, error) {
	wire.Build(
		generators.NewClientEventConsumerAwsLambdaGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventProducerGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventProducerGeneratorService, error) {
	wire.Build(
		generators.NewClientEventProducerGeneratorService,
		os.NewOsHelper)

	return nil, nil
}

func NewClientEventGenerateGeneratorService(logger *zap.SugaredLogger) (generators.ClientEventGenerateGeneratorService, error) {
	wire.Build(
		generators.NewClientEventGenerateGeneratorService,
		os.NewOsHelper)

	return nil, nil
}
