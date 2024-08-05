package database

import (
	"sync"

	databaseConfig "github.com/ilfey/hikilist-go/config/database"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/orm"
	"github.com/ilfey/hikilist-go/internal/orm/drivers"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var (
	instance orm.DB
	once     sync.Once
)

func New(config *databaseConfig.Config) orm.DB {
	once.Do(func() {
		db := sqlx.MustConnect("sqlite3", config.DBName)

		if err := db.Ping(); err != nil {
			logger.Fatalf("Database connection failed: %v", err)
		}

		instance = drivers.NewSQLX(db)
	})

	return instance
}

func Instance() orm.DB {
	if instance == nil {
		logger.Fatal("Database is not initialized")
	}

	return instance
}
