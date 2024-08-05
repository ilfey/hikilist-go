package sql

import (
	"context"
	"fmt"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/orm/database"
)

type DeleteStruct struct {
	table string
	where string
}

func (d *DeleteStruct) Where(conds any) *DeleteStruct {
	d.where = where(d.table, conds)

	// Self-return
	return d
}

func (d *DeleteStruct) SQL() string {
	sql := fmt.Sprintf("DELETE FROM %s", d.table)

	if d.where != "" {
		sql += fmt.Sprintf(" WHERE %s", d.where)
	}

	sql += ";"

	logger.Tracef("SQL: %s", sql)

	return sql
}

func (d *DeleteStruct) Exec(ctx context.Context, db database.DB) (uint, error) {
	result, err := db.Exec(ctx, d.SQL())
	if err != nil {
		return 0, err
	}

	id := errorsx.Must(result.LastInsertId())

	return uint(id), nil
}

func DeleteFromStruct[T any](data *T, table ...string) *DeleteStruct {
	var tableName string

	if len(table) != 0 {
		tableName = table[0]
	} else {
		var anyCast any = data
		if model, ok := anyCast.(TableName); ok {
			tableName = model.TableName()
		} else {
			panic("table name not found")
		}
	}

	return &DeleteStruct{
		table: tableName,
	}
}
