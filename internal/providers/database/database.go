package database

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/config/database"
	postgres2 "github.com/ilfey/hikilist-go/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(config *database.Config) (postgres2.DB, error) {
	squirrel.StatementBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	pool, err := pgxpool.New(context.Background(), config.DSN())
	if err != nil {
		return nil, err
	}

	return &postgres2.ConnectionPool{
		Pool: pool,
	}, nil
}
