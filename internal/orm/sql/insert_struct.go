package sql

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ilfey/hikilist-go/internal/errorsx"
	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/orm/database"
)

type InsertStruct struct {
	table  string
	fields []*StructField
}

func (i *InsertStruct) Ignore(fields ...string) *InsertStruct {
	for _, field := range i.fields {
		if field.ignore {
			continue
		}

		for _, ignoreField := range fields {
			if field.Name() == ignoreField || field.Snake() == ignoreField {
				field.ignore = true
			}
		}
	}

	return i
}

func (i *InsertStruct) queryColumns() string {
	expr := ""

	for _, field := range i.fields {
		if field.ignore {
			continue
		}

		expr += fmt.Sprintf("%s, ", field.Snake())
	}

	return strings.TrimSuffix(expr, ", ")
}

func (i *InsertStruct) queryValues() string {
	expr := ""

	for _, field := range i.fields {
		if field.ignore {
			continue
		}

		expr += fmt.Sprintf("%s, ", field.ValueString())
	}

	return strings.TrimSuffix(expr, ", ")
}

func (i *InsertStruct) SQL() string {
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		i.table,
		i.queryColumns(),
		i.queryValues(),
	)

	logger.Tracef("SQL: %s", sql)

	return sql

}

func (i *InsertStruct) Clean() *InsertStruct {
	for _, field := range i.fields {
		field.ignore = false
	}

	// Self-return
	return i
}

func (i *InsertStruct) Exec(ctx context.Context, db database.DB) (uint, error) {
	result, err := db.Exec(ctx, i.SQL())
	if err != nil {
		return 0, err
	}

	id := errorsx.Must(result.LastInsertId())

	return uint(id), nil
}

func InsertFromStruct[T any](data *T, table ...string) *InsertStruct {
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

	value := reflect.ValueOf(data).Elem()

	if value.Kind() != reflect.Struct {
		panic("dest must be struct")
	}

	is := InsertStruct{
		table:  tableName,
		fields: make([]*StructField, value.NumField()),
	}

	for i := 0; i < value.NumField(); i++ {
		is.fields[i] = &StructField{
			reflectValue: value.Field(i),
			name:         value.Type().Field(i).Name,
			idx:          i,
		}
	}

	return &is
}
