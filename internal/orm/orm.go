package orm

import (
	"github.com/ilfey/hikilist-go/internal/orm/database"
	"github.com/ilfey/hikilist-go/internal/orm/sql"
)

type DB interface {
	database.DB
}

func New(db database.DB) DB {
	return db
}

func Select[T any](dest *T, table ...string) *sql.SelectStruct[T] {
	return sql.SelectFromStruct(dest, table...)
}

func Insert[T any](dest *T, table ...string) *sql.InsertStruct {
	return sql.InsertFromStruct(dest, table...)
}

func Update[T any](dest *T, table ...string) *sql.UpdateStruct {
	return sql.UpdateFromStruct(dest, table...)
}

func Delete[T any](dest *T, table ...string) *sql.DeleteStruct {
	return sql.DeleteFromStruct(dest, table...)
}
