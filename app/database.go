package app

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB = &bun.DB{}

func InitDB() (_ *bun.DB, err error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		Config.DbUser,
		Config.DbPassword,
		Config.DbHost,
		Config.DbPort,
		Config.DbName,
	)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(conn, pgdialect.New())

	if Config.Debug {
		DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))
	}

	return DB, nil
}
