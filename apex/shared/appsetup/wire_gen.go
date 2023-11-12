// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/sigmasee/sigmasee/apex/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewEntgoClient(logger *zap.SugaredLogger, databaseConfig database.DatabaseConfig, postgresConfig postgres.PostgresConfig, appConfig configuration.AppConfig) (repositories.EntgoClient, error) {
	databaseDatabase, err := postgres.NewPostgres(logger, appConfig, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := repositories.NewEntgoClient(logger, databaseConfig, databaseDatabase)
	if err != nil {
		return nil, err
	}
	return entgoClient, nil
}