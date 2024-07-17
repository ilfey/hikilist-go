package database

import (
	databaseConfig "github.com/ilfey/hikilist-go/config/database"
	"github.com/ilfey/hikilist-go/internal/errorsx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(config *databaseConfig.Config) *gorm.DB {
	return errorsx.Must(
		gorm.Open(
			sqlite.Open(config.DBName),
			&gorm.Config{},
		),
	)
}
