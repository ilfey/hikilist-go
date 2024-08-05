package sql

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/orm/database"
	"github.com/rotisserie/eris"
)

type SelectStruct[T any] struct {
	model *T

	fields    []*StructField
	resolvers []func(context.Context, *T) error

	table string
	alias string

	where  string
	group  string
	order  string
	offset int
	limit  int
}

func (s *SelectStruct[T]) tableOrAlias() string {
	if s.alias != "" {
		return s.alias
	}

	return s.table
}

func (s *SelectStruct[T]) As(alias string) *SelectStruct[T] {
	s.alias = alias

	// Self-return
	return s
}

func (s *SelectStruct[T]) Where(conds any) *SelectStruct[T] {
	s.where = where(s.tableOrAlias(), conds)

	// Self-return
	return s
}

func (s *SelectStruct[T]) Group(g string) *SelectStruct[T] {
	s.group = g

	// Self-return
	return s
}

func (s *SelectStruct[T]) Order(o string) *SelectStruct[T] {
	s.order = o

	// Self-return
	return s
}

func (s *SelectStruct[T]) Offset(o int) *SelectStruct[T] {
	s.offset = o

	// Self-return
	return s
}

func (s *SelectStruct[T]) Limit(l int) *SelectStruct[T] {
	s.limit = l

	// Self-return
	return s
}

func (s *SelectStruct[T]) queryExpression() string {
	expr := ""

	for _, field := range s.fields {
		if field.ignore {
			continue
		}

		expr += fmt.Sprintf("%s, ", field.GetColumn(s.tableOrAlias()))
	}

	return strings.TrimSuffix(expr, ", ")
}

func (s *SelectStruct[T]) Ignore(fields ...string) *SelectStruct[T] {
	for _, field := range s.fields {
		if field.ignore {
			continue
		}

		for _, ignoreField := range fields {
			if field.Name() == ignoreField || field.Snake() == ignoreField {
				field.ignore = true
			}
		}
	}

	return s
}

func (s *SelectStruct[T]) Resolve(field string, fn func(context.Context, *T) error) *SelectStruct[T] {
	hasField := false

	// Find field index
	for _, f := range s.fields {
		if f.Name() == field || f.Snake() == field {
			if f.ignore {
				panic(fmt.Sprintf("field %s is ignored", field))
			}

			hasField = true
			f.ignore = true
			break
		}
	}

	if !hasField {
		panic(fmt.Sprintf("field %s not found", field))
	}

	// Add resolver
	s.resolvers = append(s.resolvers, fn)

	return s
}

func (s *SelectStruct[T]) SQL() string {
	sql := fmt.Sprintf("SELECT %s FROM %s", s.queryExpression(), s.table)

	if s.alias != "" {
		sql += fmt.Sprintf(" AS %s", s.alias)
	}

	if s.where != "" {
		sql += fmt.Sprintf(" WHERE %s", s.where)
	}

	if s.group != "" {
		sql += fmt.Sprintf(" GROUP BY %s", s.group)
	}

	if s.order != "" {
		sql += fmt.Sprintf(" ORDER BY %s", s.order)
	}

	if s.offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", s.offset)
	}

	if s.limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", s.limit)
	}

	sql += ";"

	logger.Tracef("SQL: %s", sql)

	return sql
}

func (s *SelectStruct[T]) ScanRow(row interface{ Scan(...any) error }, dest *T) error {
	value := reflect.ValueOf(dest).Elem()

	// Collect field ptrs
	var ptrs []any

	for _, field := range s.fields {
		if field.ignore {
			continue
		}

		if value.CanAddr() {
			ptrs = append(ptrs, value.Field(field.idx).Addr().Interface())
		} else {
			ptrs = append(ptrs, value.Field(field.idx).Interface())
		}
	}

	err := row.Scan(ptrs...)
	if err != nil {
		return err
	}

	return nil
}

func (s *SelectStruct[T]) QueryRow(ctx context.Context, db database.DB) (*T, error) {
	return s.QueryRowSQL(ctx, db, s.SQL())
}

func (s *SelectStruct[T]) QueryRowSQL(ctx context.Context, db database.DB, sql string) (*T, error) {
	var data T

	row := db.QueryRow(ctx, sql)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := s.ScanRow(row, &data)
	if err != nil {
		return nil, err
	}

	for _, resolver := range s.resolvers {
		err = resolver(ctx, &data)
		if err != nil {
			return nil, eris.Wrap(err, "failed to resolve")
		}
	}

	return &data, nil
}

func (s *SelectStruct[T]) Query(ctx context.Context, db database.DB) ([]*T, error) {
	return s.QuerySQL(ctx, db, s.SQL())
}

func (s *SelectStruct[T]) QuerySQL(ctx context.Context, db database.DB, sql string) ([]*T, error) {
	rows, err := db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data []*T
	for rows.Next() {
		item := new(T)
		err := s.ScanRow(rows, item)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}

func SelectFromStruct[T any](model *T, table ...string) *SelectStruct[T] {
	var tableName string

	if len(table) != 0 {
		tableName = table[0]
	} else {
		var anyCast any = model
		if model, ok := anyCast.(TableName); ok {
			tableName = model.TableName()
		} else {
			panic("table name not found")
		}
	}

	value := reflect.ValueOf(model)

	// This should never happen because always be a pointer
	// if value.Kind() != reflect.Pointer {
	// 	panic("dest must be pointer")
	// }

	value = value.Elem()

	sel := SelectStruct[T]{
		model:  model,
		table:  tableName,
		fields: make([]*StructField, value.NumField()),
	}

	for i := 0; i < value.NumField(); i++ {
		sel.fields[i] = &StructField{
			reflectValue: value.Field(i),
			name:         value.Type().Field(i).Name,
			idx:          i,
		}
	}

	return &sel
}
