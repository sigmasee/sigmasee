package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type postgres struct {
	logger *zap.SugaredLogger
	db     *sql.DB
}

func NewPostgres(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	postgresConfig PostgresConfig) (database.Database, error) {

	connectionString := postgresConfig.ConnectionString
	connectionString = strings.TrimRight(connectionString, " ?")
	source := strings.Replace(appConfig.GetSource(), "::", "_", -1)

	if strings.Contains(connectionString, "?") {
		connectionString = fmt.Sprintf("%s&application_name=%s", connectionString, source)
	} else {
		connectionString = fmt.Sprintf("%s?application_name=%s", connectionString, source)
	}

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(postgresConfig.MaxOpenConnections)

	return &postgres{
		db:     db,
		logger: logger,
	}, nil
}

func (s *postgres) GetDriver() *entsql.Driver {
	return entsql.OpenDB(dialect.Postgres, s.db)
}

func (s *postgres) GetDB() *sql.DB {
	return s.db
}

func (s *postgres) Close() {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			s.logger.Errorf("Failed to close database. Error: %v", err)
		}

		s.db = nil
	}
}
