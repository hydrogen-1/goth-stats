package database

import (
	"context"
	"errors"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/hydrogen-1/goth-stats/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(context context.Context, config *config.Config, migrations fs.FS) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:5432/gothstats?sslmode=disable", config.DbUser, config.DbPassword, config.DbHost)

	pool, err := pgxpool.New(context, connectionString)
	if err != nil {
		return nil, err
	}

	fs, err := iofs.New(migrations, "sql/migrations")
	if err != nil {
		return nil, err
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", fs, connectionString)
	if err != nil {
		return nil, err
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return pool, nil
}
