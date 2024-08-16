package database

import (
	"context"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/postgres"
	databaseConfig "github.com/ilfey/hikilist-go/pkg/config/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	instance postgres.DB
	once     sync.Once
)

func New(config *databaseConfig.Config) postgres.DB {
	once.Do(func() {
		sq.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		pool, err := pgxpool.New(context.Background(), config.DSN())
		if err != nil {
			logger.Fatalf("Database connection failed %v", err)
		}

		instance = &postgres.ConnectionPool{
			Pool: pool,
		}
	})

	return instance
}

func Instance() postgres.DB {
	if instance == nil {
		logger.Fatal("Database is not initialized")
	}

	return instance
}
