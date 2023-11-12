package database

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
)

type Database interface {
	GetDB() *sql.DB
	GetDriver() *entsql.Driver
	Close()
}
