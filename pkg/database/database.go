package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/postgres"
	databaseConfig "github.com/ilfey/hikilist-go/pkg/config/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(config *databaseConfig.Config) postgres.DB {
	sq.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	pool, err := pgxpool.New(context.Background(), config.DSN())
	if err != nil {
		logger.Fatalf("Database connection failed %v", err)
	}

	return &postgres.ConnectionPool{
		Pool: pool,
	}
}
