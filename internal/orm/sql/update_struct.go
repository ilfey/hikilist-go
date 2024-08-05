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

type UpdateStruct struct {
	fields []*StructField
	table  string
	where  string
}

func (u *UpdateStruct) Where(conds any) *UpdateStruct {
	u.where = where(u.table, conds)

	// Self-return
	return u
}

func (u *UpdateStruct) Clean() *UpdateStruct {
	u.where = ""

	for _, field := range u.fields {
		field.ignore = false
	}

	// Self-return
	return u
}

func (u *UpdateStruct) Ignore(fields ...string) *UpdateStruct {
	for _, field := range u.fields {
		if field.ignore {
			continue
		}

		for _, ignoreField := range fields {
			if field.Name() == ignoreField || field.Snake() == ignoreField {
				field.ignore = true
			}
		}
	}

	return u
}

func (u *UpdateStruct) IgnoreZeros() *UpdateStruct {
	for _, field := range u.fields {
		if field.ignore {
			continue
		}

		if field.ReflectValue().IsZero() {
			field.ignore = true
		}
	}

	// Self-return
	return u
}

func (u *UpdateStruct) querySet() string {
	expr := ""

	for _, field := range u.fields {
		if field.ignore {
			continue
		}

		expr += fmt.Sprintf("%s = %s, ", field.GetColumn(u.table), field.ValueString())
	}

	return strings.TrimSuffix(expr, ", ")
}

func (u *UpdateStruct) SQL() string {
	sql := fmt.Sprintf("UPDATE %s SET %s", u.table, u.querySet())

	if u.where != "" {
		sql += fmt.Sprintf(" WHERE %s", u.where)
	}

	sql += ";"

	logger.Tracef("SQL: %s", sql)

	return sql
}

func (u *UpdateStruct) Exec(ctx context.Context, db database.DB) (uint, error) {
	result, err := db.Exec(ctx, u.SQL())
	if err != nil {
		return 0, err
	}

	id := errorsx.Must(result.LastInsertId())

	return uint(id), nil
}

func UpdateFromStruct[T any](data *T, table ...string) *UpdateStruct {
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

	us := UpdateStruct{
		table:  tableName,
		fields: make([]*StructField, value.NumField()),
	}

	for i := 0; i < value.NumField(); i++ {
		us.fields[i] = &StructField{
			reflectValue: value.Field(i),
			name:         value.Type().Field(i).Name,
			idx:          i,
		}
	}

	return &us
}
